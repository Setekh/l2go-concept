package server

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
)

func WritePlayOk(sessionId uint64, key auth.SessionKey) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x07)
	buffer.WriteUInt32(uint32(key.PlayOk1))
	buffer.WriteUInt32(uint32(key.PlayOk2))
	//buffer.WriteUInt64(sessionId)

	return buffer
}
