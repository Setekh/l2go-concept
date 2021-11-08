package model

import (
	"gorm.io/gorm"
	"time"
)

type Location struct {
	X       int32
	Y       int32
	Z       int32
	Heading int32
}

type CoreStats struct {
	INT uint32
	STR uint32
	CON uint32
	MEN uint32
	DEX uint32
	WIT uint32
}

type HealthManaStats struct {
	CurrentHp float64
	MaxHp     float64 `gorm:"-"`
	CurrentMp float64
	MaxMp     float64 `gorm:"-"`
}

const (
	ShortcutTypeItem   = 1
	ShortcutTypeSkill  = 2
	ShortcutTypeAction = 3
	ShortcutTypeMacro  = 4
	ShortcutTypeRecipe = 5
)

type UserShortcut struct {
	Id    uint32
	Level int
	Slot  uint32
	Page  uint32
	Type  uint32
}

type Npc struct {
	EntityId uint32
	Name     string
	Title    string
	Level    uint32
	Location
	CoreStats
}

type Character struct {
	gorm.Model

	EntityId    uint32 `gorm:"-"`
	AccountName string `gorm:"not null"`
	Name        string `gorm:"uniqueIndex"`
	Title       string
	Level       uint32
	Location
	CoreStats   `gorm:"-"`
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
	HealthManaStats
	AccessLevel  uint32
	LastAccessed time.Time
	MaxCp        uint32 `gorm:"-"`
	CurrentCp    uint32
	Hero         bool
	PvpKills     uint32
	PkKills      uint32
	Shortcuts    []UserShortcut `gorm:"-"`
	Destination  *Location      `gorm:"-"`
}
