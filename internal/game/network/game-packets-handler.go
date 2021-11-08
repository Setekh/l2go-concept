package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/network/clientPackets"
	"l2go-concept/pkg/game/client"
)

var packets = map[uint]client.IncomingPacket{}

func init() {
	packets[0x00] = &clientPackets.RequestKeys{}
	packets[0x08] = &clientPackets.RequestCharacterList{}
	packets[0x0b] = &clientPackets.RequestCreateCharacter{}
	packets[0x0e] = &clientPackets.RequestCharacterCreateScreen{}
	packets[0x0d] = &clientPackets.RequestSelectCharacter{}
	packets[0x09] = &clientPackets.RequestLogout{}
	packets[0x46] = &clientPackets.RequestRestart{}
	packets[0x01] = &clientPackets.RequestMove{}
}

func HandlePacket(client *Client, dm game.DependencyManager, opcode uint, bytes []byte) {
	var reader = common.NewReader(bytes)

	packet := packets[opcode]
	if packet != nil {
		packet.ReadPacket(client, dm, reader)
		return
	}

	switch opcode {
	case 0x03:
		{
			list := []model.Skill{
				{
					Id:      3,
					Level:   9,
					Passive: false,
				},
			}

			client.SendRawPacket(SkillList(list))
			client.SendRawPacket(FriendList())
			client.SendRawPacket(QuestList())
			client.SendRawPacket(HennaInfo(client.player))
			client.SendRawPacket(ClientSetTime(dm.GetTimeController()))
			client.SendRawPacket(UserShortcutInit(client.player))
			client.SendRawPacket(UserItemList(client.player, false))
			client.SendRawPacket(ValidateLocation(client.player))
			client.SendRawPacket(UserInfo(client))
			client.SendRawPacket(ActionFailed())
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

			client.SendRawPacket(SkillList(list))
			break
		}
	}
}
