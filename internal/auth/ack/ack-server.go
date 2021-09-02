package ack

import (
	"l2go-concept/internal/network"
)

func ClientAckPacket(blowKey []byte, modulus []byte) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x00)
	buffer.Write([]byte{0xfd, 0x8a, 0x22, 0x00}) // Session id?
	buffer.Write([]byte{0x5a, 0x78, 0x00, 0x00}) // Protocol version : 785a

	buffer.Write(modulus) // RSA Public

	// unk GG related?
	buffer.WriteUInt32(0x29DD954E)
	buffer.WriteUInt32(0x77C39CFC)
	buffer.WriteUInt32(0x97ADB620)
	buffer.WriteUInt32(0x07BDE0F7)

	buffer.Write(blowKey)
	buffer.WriteByte(0x00) // lol

	return buffer
}
