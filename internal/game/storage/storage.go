package storage

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"l2go-concept/internal/auth/model"
	"log"
)

type GameStorage interface {
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

	err = db.AutoMigrate(&model.Account{})
	if err != nil {
		log.Fatalln("Failed db migration!", err)
	}

	return &context{
		db: db,
	}
}
