package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"l2go-concept/internal/game/model"
	"log"
)

type GameStorage interface {
	StoreNewCharacter(character *model.Character)
	LoadAllCharacters(accountName string) []model.Character
	CheckNameExists(name string) bool
	LoadCharacter(name string, slot uint32) *model.Character
	SaveCharacter(player *model.Character)
}

type context struct {
	db *gorm.DB
}

func CreateStorage() GameStorage {
	open := sqlite.Open("game.db")
	db, err := gorm.Open(open, &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed db connection!", err)
	}

	err = db.AutoMigrate(&model.Character{})
	if err != nil {
		log.Fatalln("Failed db migration!", err)
	}

	return &context{
		db: db,
	}
}

func (store context) StoreNewCharacter(character *model.Character) {
	tx := store.db.Create(&character)
	if tx.Error != nil {
		log.Printf("Failed storing character %+v. Caused by %s", character, tx.Error.Error())
	}
}

func (store *context) LoadAllCharacters(accountName string) []model.Character {
	var list []model.Character
	store.db.Where("account_name = ?", accountName).Order("created_at asc").Find(&list)

	for index, _ := range list { // TODO implement templates
		list[index].MaxHp = 1200
		list[index].MaxMp = 1200

		list[index].INT = 21
		list[index].STR = 22
		list[index].MEN = 23
		list[index].CON = 24
		list[index].DEX = 25
		list[index].WIT = 26

		list[index].X = -71338
		list[index].Y = 258271
		list[index].Z = -3104
		list[index].Heading = 56277
	}

	return list
}

func (store *context) SaveCharacter(player *model.Character) {
	store.db.Updates(player)
}

func (store *context) LoadCharacter(accountName string, slot uint32) *model.Character {
	var list = store.LoadAllCharacters(accountName)
	return &list[slot]
}

func (store *context) CheckNameExists(name string) bool {
	var count int64
	store.db.Model(&model.Character{}).Where("name = ?", name).Count(&count)
	return count > 0
}
