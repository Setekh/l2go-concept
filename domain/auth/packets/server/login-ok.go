package server

import (
	"l2go-concept/domain/network"
)

func WriteLoginOk(sessionId uint64) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x03)
	buffer.WriteUInt64(sessionId)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x000003ea)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.Write(make([]byte, 16))
	return buffer
}
