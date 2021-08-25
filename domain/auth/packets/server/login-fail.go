package server

import (
	"l2go-concept/domain/network"
)

const (
	_NONE = iota
	SystemError
	AccountPasswordWrong
	AccountOrPasswordWrong
	AccessFailed
	AccountInUse = 0x07
)

func WriteLoginFail(reason uint32) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x06)
	buffer.WriteUInt32(reason)
	return buffer
}
