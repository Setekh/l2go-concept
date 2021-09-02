package gameguard

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
)

type GGPacketHandler struct {
}

func (h *GGPacketHandler) HandlePacket(_ uint, client auth.Client, reader *network.Reader, _ auth.Storage) {
	authGG := RequestGGAuth(reader)

	log.Println("GG trys to authenticate!", authGG.SessionId)

	if client == nil {
		return
	}

	client.SendPacketEncoded(GGAuthResponse(authGG.SessionId))
}
