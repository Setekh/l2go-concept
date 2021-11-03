package network

import (
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
)

var packets = map[uint]auth.ClientPacket{}

func init() {
	packets[0x00] = &RequestAuth{}
	packets[0x02] = &RequestPlayServer{}
	packets[0x05] = &RequestServerList{}
	packets[0x07] = &RequestGGAuth{}
}

func HandlePacket(client *Client, store auth.Storage, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	handler := packets[opcode]

	if handler != nil {
		// TODO check here if we f*ed the performance somewhat
		ctx := auth.Context{
			Client:  client,
			Storage: store,
		}

		handler.HandlePacket(reader, ctx)
	} else {
		log.Printf("Failed fetching packet handler for packet %X\n", opcode)
	}
}
