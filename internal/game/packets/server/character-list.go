package server

import (
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/network"
)

func WriteCharacterList(characters []model.Character) *network.Buffer {
	buffer := network.NewBuffer()
	buffer.WriteByte(0x13)   // Packet type: CharList
	buffer.WriteUInt32(0x00) // TODO

	return buffer
}
