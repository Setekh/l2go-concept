package clientPackets

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/network/crypt"
	game2 "l2go-concept/pkg/game/client"
	"log"
)

type RequestKeys struct{}

func (r *RequestKeys) ReadPacket(client game2.L2Client, dm game.DependencyManager, reader *common.Reader) {
	var protocolVersion = reader.ReadD()
	log.Printf("L2Client is with protocol version %d\n", protocolVersion)

	client.SendPacket(&SendCryptKeys{}, crypt.GetKey())
	client.EnableCrypt()
}

type SendCryptKeys struct{}

func (s *SendCryptKeys) WritePacket(buffer *common.Buffer, params ...interface{}) {
	cryptKey := (params[0]).([]byte)

	buffer.WriteC(0x00)
	buffer.WriteC(0x01) // Protocol is ok
	buffer.WriteBytes(cryptKey)
	buffer.WriteD(0x01) // server id ?
	buffer.WriteD(0x01)
}
