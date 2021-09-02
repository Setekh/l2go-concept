package model

import "gorm.io/gorm"

type GameServer struct {
	gorm.Model

	ServerId       uint8
	Ip             string
	Port           uint32
	CurrentPlayers uint32
	MaxPlayers     uint32
	PvpServer      bool
	IsUp           bool
	ServerType     uint32
	ServerBrackets bool
}
