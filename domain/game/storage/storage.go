package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"l2go-concept/domain/auth/model"
	"log"
)

type GameStorage interface {
}

type context struct {
	db *gorm.DB
}

func CreateStorage() GameStorage {
	db, err := gorm.Open(sqlite.Open("game.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed db connection!", err)
	}

	err = db.AutoMigrate(&model.Account{})
	if err != nil {
		log.Fatalln("Failed db migration!", err)
	}

	return &context{
		db: db,
	}
}
