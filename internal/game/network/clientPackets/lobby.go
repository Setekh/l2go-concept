package clientPackets

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
	"l2go-concept/pkg/game/client"
	"time"
)

type RequestCharacterList struct{}

func (r *RequestCharacterList) ReadPacket(client client.L2Client, dm game.DependencyManager, reader *common.Reader) {
	accountName := reader.ReadString()
	playOk2 := reader.ReadD()
	playOk1 := reader.ReadD()
	loginOk1 := reader.ReadD()
	loginOk2 := reader.ReadD()

	println(accountName, loginOk1, loginOk2, playOk1, playOk2)

	client.Upgrade(accountName, playOk1)

	store := dm.GetStorage()
	characters := store.LoadAllCharacters(accountName)

	client.SendPacket(&CharacterList{
		characters,
		client.GetAccountName(),
		client.GetSessionId(),
	})
}

type CharacterList struct {
	Characters  []model.Character
	AccountName string
	SessionId   uint32
}

func (c *CharacterList) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	characters := c.Characters
	accountName := c.AccountName
	sessionId := c.SessionId

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

		buffer.WriteS(accountName)
		buffer.WriteD(sessionId)

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
}

type RequestSelectCharacter struct{}

func (r *RequestSelectCharacter) ReadPacket(client client.L2Client, dm game.DependencyManager, reader *common.Reader) {
	store := dm.GetStorage()
	slot := reader.ReadD()

	var character = store.LoadCharacter(client.GetAccountName(), slot)
	if character == nil || character.AccessLevel < 0 {
		client.Close()
		return
	}

	character.EntityId = slot + 1 // TODO Please no
	character.LastAccessed = time.Now()
	store.SaveCharacter(character)

	client.SetPlayer(character)

	client.SendPacket(&SelectedCharacter{
		Character: character,
		SessionId: client.GetSessionId(),
		GameTime:  uint32(dm.GetTimeController().GetGameTime()),
	})
}

type SelectedCharacter struct {
	Character *model.Character
	SessionId uint32
	GameTime  uint32
}

func (s *SelectedCharacter) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	character := s.Character
	sessionId := s.SessionId
	gameTime := s.GameTime

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
	buffer.WriteD(gameTime)

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
