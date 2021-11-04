package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/storage"
)

func RequestCharacterList(client *Client, store storage.GameStorage, reader *common.Reader) {
	accountName := reader.ReadString()
	playOk2 := reader.ReadD()
	playOk1 := reader.ReadD()
	loginOk1 := reader.ReadD()
	loginOk2 := reader.ReadD()

	println(accountName, loginOk1, loginOk2, playOk1, playOk2)

	client.accountName = accountName
	client.playOk = playOk1

	characters := store.LoadAllCharacters(accountName)

	client.SendPacket(WriteCharacterList(client, characters))
}

func WriteCharacterList(client *Client, characters []model.Character) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0x13)
	buffer.WriteD(uint32(len(characters)))

	var lastTime = characters[0].LastAccessed
	var lastActiveId = 1
	for charId, character := range characters {
		if character.LastAccessed.After(lastTime) {
			lastTime = character.LastAccessed
			lastActiveId = charId + 1
		}

		buffer.WriteS(character.Name)
		buffer.WriteD(uint32(charId + 1)) // Todo this should be a world entity value

		buffer.WriteS(client.accountName)
		buffer.WriteD(client.playOk)

		buffer.WriteD(character.ClanId)
		buffer.WriteD(0x00) // Unk

		buffer.WriteD(character.Sex)
		buffer.WriteD(character.Race)

		buffer.WriteD(character.ClassId)
		buffer.WriteD(0x01) // Active ?

		buffer.WriteD(0x00) // X
		buffer.WriteD(0x00) // Y
		buffer.WriteD(0x00) // Z

		buffer.WriteF(character.CurrentHp)
		buffer.WriteF(character.CurrentMp)

		buffer.WriteD(character.SkillPoints)
		buffer.WriteD(character.Experience)
		buffer.WriteD(character.Level)

		buffer.WriteD(character.Karma)
		for i := 0; i < 9; i++ {
			buffer.WriteD(0x00) // Unk
		}

		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_UNDER)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_REAR)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LEAR)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_NECK)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_RFINGER)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LFINGER)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_HEAD)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_RHAND)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LHAND)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_GLOVES)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_CHEST)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LEGS)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_FEET)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_BACK)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LRHAND)
		buffer.WriteD(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_HAIR)

		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_UNDER)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_REAR)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LEAR)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_NECK)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_RFINGER)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LFINGER)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_HEAD)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_RHAND)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LHAND)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_GLOVES)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_CHEST)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LEGS)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_FEET)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_BACK)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LRHAND)
		buffer.WriteD(0x00) // PaperdollItemId(Inventory.PAPERDOLL_HAIR)

		buffer.WriteD(character.Hair)
		buffer.WriteD(character.HairColor)
		buffer.WriteD(character.Face)

		buffer.WriteF(character.MaxHp)
		buffer.WriteF(character.MaxMp)

		buffer.WriteSD(0x00) // TODO days before delete & access level, -1 == banned
		buffer.WriteD(character.ClassId)
		buffer.WriteD(uint32(lastActiveId)) // Is active character 0x01 for active
		buffer.WriteC(127)                  // Weapon enchant, min 127?
	}

	return buffer
}

func SelectCharacter(sessionId uint32, character *model.Character, buffer *common.Buffer) {
	buffer.WriteC(0x15)
	buffer.WriteS(character.Name)
	buffer.WriteD(character.EntityId)
	buffer.WriteS(character.Title)
	buffer.WriteD(sessionId)
	buffer.WriteD(character.ClanId)
	buffer.WriteD(0x00) // Unk
	buffer.WriteD(character.Sex)
	buffer.WriteD(character.Race)
	buffer.WriteD(character.ClassId)
	buffer.WriteD(0x01) // Active?
	buffer.WriteSD(character.X)
	buffer.WriteSD(character.Y)
	buffer.WriteSD(character.Z)

	buffer.WriteF(character.CurrentHp)
	buffer.WriteF(character.CurrentMp)
	buffer.WriteD(character.SkillPoints)
	buffer.WriteD(character.Experience)
	buffer.WriteD(character.Level)
	buffer.WriteD(character.Karma) // thx evill33t
	buffer.WriteD(0x0)             // ?
	buffer.WriteD(character.INT)
	buffer.WriteD(character.STR)
	buffer.WriteD(character.CON)
	buffer.WriteD(character.MEN)
	buffer.WriteD(character.DEX)
	buffer.WriteD(character.WIT)

	buffer.WriteBytes(make([]byte, 30))

	buffer.WriteD(0x00) // C3 work
	buffer.WriteD(0x00) // C3 work

	// extra info
	buffer.WriteD(0x123) // TODO in-game time

	buffer.WriteD(0x00) //

	buffer.WriteD(0x00) // c3

	buffer.WriteD(0x00) // c3 InspectorBin
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3

	buffer.WriteD(0x00) // c3 InspectorBin for 528 client
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
	buffer.WriteD(0x00) // c3
}
