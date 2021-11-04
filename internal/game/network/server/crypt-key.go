package server

import (
	"l2go-concept/internal/common"
)

func WriteKeyPacket(cryptKey []byte) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteByte(0x00)
	buffer.WriteByte(0x01) // Protocol is ok
	buffer.Write(cryptKey)
	buffer.WriteUInt32(0x01) // server id ?
	buffer.WriteUInt32(0x01)

	return buffer
}
