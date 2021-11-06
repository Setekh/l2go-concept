package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
)

func RequestMove(client *Client, reader *common.Reader) {
	targetX := reader.ReadSD()
	targetY := reader.ReadSD()
	targetZ := reader.ReadSD()

	//originX := reader.ReadD()
	//originY := reader.ReadD()
	//originZ := reader.ReadD()

	player := client.player
	player.Destination = &model.Location{
		X: targetX,
		Y: targetY,
		Z: targetZ,
	}

	client.SendPacket(MoveToLocation(client.player))

	// -.-
	player.Destination = nil
	player.X = targetX
	player.Y = targetY
	player.Z = targetZ
}

func MoveToLocation(character *model.Character) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0x01)
	buffer.WriteD(character.EntityId)
	buffer.WriteSD(character.Destination.X)
	buffer.WriteSD(character.Destination.Y)
	buffer.WriteSD(character.Destination.Z)

	buffer.WriteSD(character.X)
	buffer.WriteSD(character.Y)
	buffer.WriteSD(character.Z)
	return buffer
}
