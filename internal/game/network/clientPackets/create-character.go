package clientPackets

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
	"l2go-concept/pkg/game/client"
)

type CreateCharacterScreenOk struct{}

var CreateCharacterScreenOkStatic = &CreateCharacterScreenOk{}

func (r *CreateCharacterScreenOk) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	buffer.WriteC(0x17)
}

type RequestCharacterCreateScreen struct{}

func (r *RequestCharacterCreateScreen) ReadPacket(client client.L2Client, dm game.DependencyManager, reader *common.Reader) {
	client.SendPacket(CreateCharacterScreenOkStatic)
}

type RequestCreateCharacter struct{}

func (r *RequestCreateCharacter) ReadPacket(client client.L2Client, dm game.DependencyManager, buff *common.Reader) {
	accountName := client.GetAccountName()
	store := dm.GetStorage()

	name := buff.ReadString()

	nameAlreadyExists := store.CheckNameExists(name)
	if nameAlreadyExists {
		client.SendPacket(&characterCreateFail{
			reason: ReasonNameAlreadyExists,
		})
		return
	}

	race := buff.ReadD()
	sex := buff.ReadD()
	classId := buff.ReadD()

	//buff.ReadD() // int
	//buff.ReadD() // str
	//buff.ReadD() // con
	//buff.ReadD() // men
	//buff.ReadD() // dex
	//buff.ReadD() // wit

	hairStyle := buff.ReadD()
	hairColor := buff.ReadD()
	face := buff.ReadD()

	character := &model.Character{
		AccountName: accountName,
		Name:        name,
		Level:       1,
		SkillPoints: 5000,
		ClassId:     classId,
		Sex:         sex,
		Race:        race,
		Face:        face,
		Hair:        hairStyle,
		HairColor:   hairColor,
		HealthManaStats: model.HealthManaStats{
			CurrentHp: 80,
			MaxHp:     80,
			CurrentMp: 60,
			MaxMp:     60,
		},
	}

	store.StoreNewCharacter(character)

	client.SendPacket(&characterCreateOk{})

	characters := store.LoadAllCharacters(accountName)
	client.SendPacket(&CharacterList{
		Characters:  characters,
		AccountName: accountName,
		SessionId:   client.GetSessionId(),
	})
}

type characterCreateOk struct{}

func (c *characterCreateOk) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	buffer.WriteC(0x19)
	buffer.WriteD(0x01)
}

const (
	ReasonCreationFailed = iota
	ReasonTooManyCharacters
	ReasonNameAlreadyExists
	Reason16EngChars
	ReasonIncorrectName
)

type characterCreateFail struct {
	reason uint32
}

func (c *characterCreateFail) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	buffer.WriteC(0x1A)
	buffer.WriteD(c.reason)
}
