package server

import (
	"l2go-concept/internal/network"
)

var cryptKey = []uint8{0x94, 0x35, 0x00, 0x00, 0xa1, 0x6c, 0x54, 0x87}

func WriteKeyPacket() *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x00)
	buffer.WriteByte(0x01) // Protocol is ok
	buffer.Write(cryptKey)
	buffer.WriteUInt32(0x01)
	buffer.WriteUInt32(0x01)

	return buffer
}
