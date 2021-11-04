package storage

import (
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"l2go-concept/internal/auth/model"
	"l2go-concept/pkg/auth"
	"log"
)

const (
	AccountNotFound = iota
	InvalidPassword
	Ok
	Error
)

type Context struct {
	db *gorm.DB
}

func (ctx *Context) GetAccount(name string) *model.Account {
	var account model.Account
	ctx.db.Find(&account, "name = ?", name)

	if account.ID <= 0 {
		return nil
	}

	return &account
}

func (ctx *Context) VerifyAccount(name, password string) int {
	account := ctx.GetAccount(name)

	if account == nil {
		return AccountNotFound
	}

	var bytePassword = []byte(password)
	passwordHash, err := base64.StdEncoding.DecodeString(account.Password)

	if err != nil {
		log.Printf("Failed decoding password for user %s!", name)
		return Error
	}

	err = bcrypt.CompareHashAndPassword(passwordHash, bytePassword)

	if err != nil {
		return InvalidPassword
	}

	return Ok
}

func (ctx *Context) CreateAccount(name, password string) {
	var bytePassword = []byte(password)

	bytes, err := bcrypt.GenerateFromPassword(bytePassword, 10)

	if err != nil {
		log.Printf("Failed generating password for user %s!", name)
		return
	}

	ctx.db.Create(&model.Account{
		Name:     name,
		Password: base64.StdEncoding.EncodeToString(bytes),
	})
}

func CreateStorage() auth.Storage {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed db connection!", err)
	}

	err = db.AutoMigrate(&model.Account{})
	if err != nil {
		log.Fatalln("Failed db migration!", err)
	}

	return &Context{
		db: db,
	}
}
