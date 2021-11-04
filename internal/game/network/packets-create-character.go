package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/storage"
)

func RequestCreateCharacter(client *Client, buffer *common.Buffer) {
	buffer.WriteC(0x17)
	client.SendPacket(buffer)
}

func CreateCharacter(client *Client, store storage.GameStorage, buff *common.Reader) {
	name := buff.ReadString()

	nameAlreadyExists := store.CheckNameExists(name)
	if nameAlreadyExists {
		buffer := common.NewBuffer()
		characterCreateFail(REASON_NAME_ALREADY_EXISTS, buffer)
		client.SendPacket(buffer)
		return
	}

	race := buff.ReadD()
	sex := buff.ReadD()
	classId := buff.ReadD()

	buff.ReadD() // int
	buff.ReadD() // str
	buff.ReadD() // con
	buff.ReadD() // men
	buff.ReadD() // dex
	buff.ReadD() // wit

	hairStyle := buff.ReadD()
	hairColor := buff.ReadD()
	face := buff.ReadD()

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
