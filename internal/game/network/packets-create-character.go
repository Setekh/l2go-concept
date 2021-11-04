package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/storage"
	"log"
)

func RequestCreateCharacter(client *Client, buffer *common.Buffer) {
	buffer.WriteC(0x17)
	client.SendPacket(buffer)
}

func CreateCharacter(client *Client, store storage.GameStorage, buff *common.Reader) {
	name := buff.ReadString()
	race := buff.ReadUInt32()
	sex := buff.ReadUInt32()
	classId := buff.ReadUInt32()

	buff.ReadUInt32() // int
	buff.ReadUInt32() // str
	buff.ReadUInt32() // con
	buff.ReadUInt32() // men
	buff.ReadUInt32() // dex
	buff.ReadUInt32() // wit

	hairStyle := buff.ReadUInt32()
	hairColor := buff.ReadUInt32()
	face := buff.ReadUInt32()

	log.Printf("Request to create %s of race %d of sex %d of classId %d with hair style %d, hair color %d and face %d", name, race, sex, classId, hairStyle, hairColor, face)
	character := &model.Character{
		EntityId:    0,
		AccountName: client.accountName,
		Name:        name,
		Level:       1,
		SkillPoints: 5000,
		Experience:  0,
		ClassId:     classId,
		Sex:         sex,
		Race:        race,
		Face:        face,
		Hair:        hairStyle,
		HairColor:   hairColor,
		ClanId:      0,
		Karma:       0,
		DeleteTime:  0,
		CurrentHp:   1000,
		MaxHp:       1200,
		CurrentMp:   500,
		MaxMp:       500,
	}

	store.StoreNewCharacter(character)

	buffer := common.NewBuffer()
	characterCreateOk(buffer)
	client.SendPacket(buffer)

	characters := store.LoadAllCharacters(client.accountName)
	client.SendPacket(WriteCharacterList(client, characters))
}

func characterCreateOk(buffer *common.Buffer) {
	buffer.WriteC(0x19)
	buffer.WriteD(0x01)
}

const (
	REASON_CREATION_FAILED = iota
	REASON_TOO_MANY_CHARACTERS
	REASON_NAME_ALREADY_EXISTS
	REASON_16_ENG_CHARS
	REASON_INCORRECT_NAME
)

func characterCreateFail(reason uint32, buffer *common.Buffer) {
	buffer.WriteC(0x1A)
	buffer.WriteD(reason)
}
