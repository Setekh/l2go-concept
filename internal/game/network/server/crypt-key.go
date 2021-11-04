package server

import (
	"l2go-concept/internal/common"
)

func WriteKeyPacket(cryptKey []byte) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x00)
	buffer.WriteC(0x01) // Protocol is ok
	buffer.WriteBytes(cryptKey)
	buffer.WriteD(0x01) // server id ?
	buffer.WriteD(0x01)

	return buffer
}
