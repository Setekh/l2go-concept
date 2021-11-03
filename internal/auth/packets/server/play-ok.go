package server

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
)

func WritePlayOk(serverId byte, key auth.SessionKey) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x07)
	buffer.WriteUInt32(key.PlayOk1)
	buffer.WriteUInt32(key.PlayOk2)
	buffer.WriteByte(serverId) // other packs

	return buffer
}
