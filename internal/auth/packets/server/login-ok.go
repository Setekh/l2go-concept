package server

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
)

func WriteLoginOk(sessionId uint64, key auth.SessionKey) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x03)

	buffer.WriteUInt32(uint32(key.LoginOk1))
	buffer.WriteUInt32(uint32(key.LoginOk2))
	//buffer.WriteUInt64(sessionId)

	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x000003ea)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.Write(make([]byte, 16))
	return buffer
}
