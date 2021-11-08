package clientPackets

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/model"
	"l2go-concept/pkg/game/client"
)

type RequestMove struct{}

func (r *RequestMove) ReadPacket(client client.L2Client, _ game.DependencyManager, reader *common.Reader) {
	targetX := reader.ReadSD()
	targetY := reader.ReadSD()
	targetZ := reader.ReadSD()

	//originX := reader.ReadD()
	//originY := reader.ReadD()
	//originZ := reader.ReadD()

	player := client.GetPlayer()
	player.Destination = &model.Location{
		X: targetX,
		Y: targetY,
		Z: targetZ,
	}

	client.SendPacket(MoveToLocationStatic, player)

	// -.-
	player.Destination = nil
	player.X = targetX
	player.Y = targetY
	player.Z = targetZ
}

type MoveToLocation struct {
	character *model.Character
}

var MoveToLocationStatic = &MoveToLocation{}

func (m *MoveToLocation) WritePacket(buffer *common.Buffer, params ...interface{}) {
	var character *model.Character

	if m.character == nil {
		character = (params[0]).(*model.Character)
	} else {
		character = m.character
	}

	buffer.WriteC(0x01)
	buffer.WriteD(character.EntityId)
	buffer.WriteSD(character.Destination.X)
	buffer.WriteSD(character.Destination.Y)
	buffer.WriteSD(character.Destination.Z)

	buffer.WriteSD(character.X)
	buffer.WriteSD(character.Y)
	buffer.WriteSD(character.Z)
}
