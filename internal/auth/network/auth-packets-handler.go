package network

import (
	"l2go-concept/internal/auth/gameguard"
	"l2go-concept/internal/auth/listing"
	"l2go-concept/internal/auth/login"
	"l2go-concept/internal/auth/packets/server"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
)

var packets = map[uint]auth.ClientPacketHandler{}

func init() {
	packets[0x00] = &login.PacketHandler{}
	packets[0x05] = &listing.PacketHandler{}
	packets[0x07] = &gameguard.GGPacketHandler{}
}

func HandlePacket(client *Client, store auth.Storage, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	handler := packets[opcode]

	if handler != nil {
		handler.HandlePacket(opcode, client, reader, store)
	} else {
		log.Printf("Failed fetching packet handler for packet %X\n", opcode)
	}

	switch opcode {
	case 0x02: // Request play
		{
			println("Yoo", client.sessionId, "heh", client.sessionKey.LoginOk1, client.sessionKey.LoginOk2, client.sessionKey.PlayOk1, client.sessionKey.PlayOk2)
			client.SendPacketEncoded(server.WritePlayOk(client.sessionId, client.sessionKey))
		}
	}
}
