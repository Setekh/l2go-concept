package auth

import "l2go-concept/internal/auth/model"

type Storage interface {
	GetAccount(name string) *model.Account
	VerifyAccount(name, password string) int
	CreateAccount(name, password string)
}
