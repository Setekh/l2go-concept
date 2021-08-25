package server

import (
	"l2go-concept/domain/network"
)

func GGAuthResponse(sessionId uint32) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x0B)
	buffer.WriteUInt32(sessionId)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)

	return buffer
}
