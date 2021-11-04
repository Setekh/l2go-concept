package model

import "gorm.io/gorm"

type Character struct {
	gorm.Model

	EntityId    uint32 `gorm:"-"`
	AccountName string `gorm:"not null"`
	Name        string `gorm:"uniqueIndex"`
	Level       uint32
	SkillPoints uint32
	Experience  uint32
	ClassId     uint32
	Sex         uint32
	Race        uint32
	Face        uint32
	Hair        uint32
	HairColor   uint32
	ClanId      uint32
	Karma       uint32
	DeleteTime  uint32
	CurrentHp   float64
	MaxHp       float64 `gorm:"-"`
	CurrentMp   float64
	MaxMp       float64 `gorm:"-"`
}
