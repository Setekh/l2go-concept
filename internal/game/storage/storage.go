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
	store.db.Where("account_name = ?", accountName).Find(&list)

	for index, _ := range list {
		list[index].MaxHp = 1200
		list[index].MaxMp = 1200
	}

	return list
}
