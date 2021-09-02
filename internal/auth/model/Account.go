package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model

	Name       string
	Password   string
	LastServer uint8
}
