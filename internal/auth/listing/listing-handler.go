package listing

import (
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
)

type PacketHandler struct {
	list []*model.GameServer
}

func (p *PacketHandler) HandlePacket(_ uint, client auth.Client, reader *network.Reader, _ auth.Storage) {
	gs := &model.GameServer{
		ServerId:   0x02,
		Ip:         "192.168.100.4",
		Port:       7777,
		MaxPlayers: 1000,
		PvpServer:  true,
		IsUp:       true,
	}

	// This will be fetched from the db most probably
	p.list = []*model.GameServer{gs}

	client.SendPacketEncoded(WriteServerList(0x00, p.list))
}
