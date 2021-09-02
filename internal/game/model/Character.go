package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model

	Name      string
	Level     uint32
	ClassId   uint32
	Sex       uint32
	Race      uint32
	Face      uint32
	Hair      uint32
	HairColor uint32
}
