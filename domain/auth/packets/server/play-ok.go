package server

import "l2go-concept/domain/network"

func WritePlayOk(sessionId uint64) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x07)
	buffer.WriteUInt64(sessionId)

	return buffer
}
