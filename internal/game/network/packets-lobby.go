package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/storage"
)

func RequestCharacterList(client *Client, store storage.GameStorage, reader *common.Reader) {
	accountName := reader.ReadString()
	playOk2 := reader.ReadUInt32()
	playOk1 := reader.ReadUInt32()
	loginOk1 := reader.ReadUInt32()
	loginOk2 := reader.ReadUInt32()

	println(accountName, loginOk1, loginOk2, playOk1, playOk2)

	client.accountName = accountName
	client.playOk = playOk1

	characters := store.LoadAllCharacters(accountName)

	client.SendPacket(WriteCharacterList(client, characters))
}

func WriteCharacterList(client *Client, characters []model.Character) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteByte(0x13)
	buffer.WriteUInt32(uint32(len(characters)))

	for charId, character := range characters {
		buffer.WriteL2String(character.Name)
		buffer.WriteUInt32(uint32(charId))

		buffer.WriteL2String(client.accountName)
		buffer.WriteUInt32(client.playOk)

		buffer.WriteUInt32(character.ClanId)
		buffer.WriteUInt32(0x00) // Unk

		buffer.WriteUInt32(character.Sex)
		buffer.WriteUInt32(character.Race)

		buffer.WriteUInt32(character.ClassId)
		buffer.WriteUInt32(0x01) // Active ?

		buffer.WriteUInt32(0x00) // X
		buffer.WriteUInt32(0x00) // Y
		buffer.WriteUInt32(0x00) // Z

		buffer.WriteUInt64(character.CurrentHp)
		buffer.WriteUInt64(character.CurrentMp)

		buffer.WriteUInt32(character.SkillPoints)
		buffer.WriteUInt32(character.Experience)
		buffer.WriteUInt32(character.Level)

		buffer.WriteUInt32(character.Karma)
		for i := 0; i < 9; i++ {
			buffer.WriteUInt32(0x00) // Unk
		}

		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_UNDER)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_REAR)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LEAR)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_NECK)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_RFINGER)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LFINGER)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_HEAD)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_RHAND)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LHAND)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_GLOVES)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_CHEST)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LEGS)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_FEET)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_BACK)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_LRHAND)
		buffer.WriteUInt32(0x00) // PaperdollObjectId(Inventory.PAPERDOLL_HAIR)

		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_UNDER)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_REAR)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LEAR)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_NECK)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_RFINGER)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LFINGER)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_HEAD)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_RHAND)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LHAND)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_GLOVES)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_CHEST)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LEGS)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_FEET)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_BACK)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_LRHAND)
		buffer.WriteUInt32(0x00) // PaperdollItemId(Inventory.PAPERDOLL_HAIR)

		buffer.WriteUInt32(character.Hair)
		buffer.WriteUInt32(character.HairColor)
		buffer.WriteUInt32(character.Face)

		buffer.WriteUInt64(character.MaxHp)
		buffer.WriteUInt64(character.MaxMp)

		buffer.WriteInt32(0x00) // days before delete & access level, -1 == banned
		buffer.WriteUInt32(character.ClassId)
		buffer.WriteUInt32(0x00) // Is active character 0x01 for active
		buffer.WriteByte(127)    // Weapon enchant, min 127?
	}

	return buffer
}
