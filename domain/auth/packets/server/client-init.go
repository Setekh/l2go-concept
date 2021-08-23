package server

import (
	"l2go-concept/domain/packets"
)

func CreateInitPacket(blowKey []byte, modulus []byte, buffer *packets.Buffer) {
	buffer.WriteByte(0x00)
	buffer.Write([]byte{0x9c, 0x77, 0xed, 0x03}) // Session id?
	buffer.Write([]byte{0x5a, 0x78, 0x00, 0x00}) // Protocol version : 785a

	buffer.Write(modulus) // RSA Public

	// unk GG related?
	buffer.WriteUInt32(0x29DD954E)
	buffer.WriteUInt32(0x77C39CFC)
	buffer.WriteUInt32(0x97ADB620)
	buffer.WriteUInt32(0x07BDE0F7)

	buffer.Write(blowKey)
	buffer.WriteByte(0x00) // lol

}
