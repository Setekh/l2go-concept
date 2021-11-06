package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/network/crypt"
	"l2go-concept/internal/game/network/server"
	"log"
	"time"
)

func HandlePacket(client *Client, dm game.DependencyManager, opcode uint, bytes []byte) {
	store := dm.GetStorage()

	var reader = common.NewReader(bytes)

	switch opcode {
	case 0x00: // Protocol
		{
			var protocolVersion = reader.ReadD()
			log.Printf("Client is with protocol version %d\n", protocolVersion)

			client.SendPacket(server.WriteKeyPacket(crypt.GetKey()))
			client.cryptEnabled = true
			break
		}
	case 0x08: // Request auth
		{
			RequestCharacterList(client, store, reader)
			break
		}
	case 0x09: // Logout
		{
			buffer := common.NewBuffer()
			buffer.WriteC(0x7e)
			client.SendPacket(buffer)
			//client.conn.Close()
			break
		}
	case 0x46: // Restart
		{
			buffer := common.NewBuffer()
			buffer.WriteC(0x5F)
			buffer.WriteD(0x01) // 1 = success
			client.SendPacket(buffer)

			characters := store.LoadAllCharacters(client.accountName)
			client.SendPacket(WriteCharacterList(client, characters))
			break
		}
	case 0x0e: // Create new Character
		{
			buffer := common.NewBuffer()
			RequestCreateCharacter(client, buffer)
			break
		}
	case 0x0b: // Request Create Character
		{
			CreateCharacter(client, store, common.NewReader(bytes))
			break
		}
	case 0x0d: // Character selected
		{
			slot := reader.ReadD()
			var character = store.LoadCharacter(client.accountName, slot)
			if character == nil || character.AccessLevel < 0 {
				client.Close()
				return
			}

			client.player = character
			client.player.EntityId = slot + 1 // TODO Please no
			client.player.LastAccessed = time.Now()
			store.SaveCharacter(client.player)

			buffer := common.NewBuffer()
			SelectCharacter(client.playOk, dm.GetTimeController().GetGameTime(), client.player, buffer)
			client.SendPacket(buffer)
			break
		}
	case 0x03:
		{
			list := []model.Skill{
				{
					Id:      3,
					Level:   9,
					Passive: false,
				},
			}

			client.SendPacket(SkillList(list))
			client.SendPacket(FriendList())
			client.SendPacket(QuestList())
			client.SendPacket(HennaInfo(client.player))
			client.SendPacket(ClientSetTime(dm.GetTimeController()))
			client.SendPacket(UserShortcutInit(client.player))
			client.SendPacket(UserItemList(client.player, false))
			client.SendPacket(ValidateLocation(client.player))
			client.SendPacket(UserInfo(client))
			client.SendPacket(ActionFailed())
			break
		}
	case 0x3f:
		{
			list := []model.Skill{
				{
					Id:      3,
					Level:   9,
					Passive: false,
				},
			}

			client.SendPacket(SkillList(list))
			break
		}
	}
}
