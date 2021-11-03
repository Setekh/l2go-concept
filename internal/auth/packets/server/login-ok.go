package server

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
)

func WriteLoginOk(key auth.SessionKey) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x03)

	buffer.WriteUInt32(key.LoginOk1)
	buffer.WriteUInt32(key.LoginOk2)

	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)

	buffer.WriteUInt32(0x000003ea) // billing type: 1002 Free, x200 paid time, x500 flat rate pre paid, others subscription
	buffer.WriteUInt32(0x00)       // paid time
	buffer.WriteUInt32(0x00)
	//buffer.WriteUInt32(0x02) - mobius??!

	buffer.WriteUInt32(0x00)       // warning mask
	buffer.Write(make([]byte, 16)) // forbidden servers
	//buffer.WriteUInt32(0x00) - l2jorg
	return buffer
}
