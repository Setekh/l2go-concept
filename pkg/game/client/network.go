package client

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
)

type OutgoingPacket interface {
	WritePacket(buffer *common.Buffer, params ...interface{})
}

type IncomingPacket interface {
	ReadPacket(client L2Client, dm game.DependencyManager, reader *common.Reader)
}

type L2Client interface {
	SendPacket(packet OutgoingPacket, params ...interface{})
	//	SendRawPacket(srcBuff *common.Buffer) error
	GetSessionId() uint32
	GetPlayer() *model.Character
	SetPlayer(character *model.Character)
	GetAccountName() string
	EnableCrypt()
	Upgrade(accountName string, sessionId uint32)
	Close()
}
