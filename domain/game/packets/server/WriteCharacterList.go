package server

import (
	"l2go-concept/domain/game/model"
	"l2go-concept/domain/network"
)

func WriteCharacterList(characters []model.Character) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x1f)                       // Packet type: CharList
	buffer.Write([]byte{0x00, 0x00, 0x00, 0x00}) // TODO

	return buffer
}
