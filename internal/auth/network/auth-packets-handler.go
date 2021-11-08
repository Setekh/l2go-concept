package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/pkg/auth"
	"l2go-concept/pkg/auth/client"
	"log"
)

var packets = map[uint]client.IncomingPacket{}

func init() {
	packets[0x00] = &RequestAuth{}
	packets[0x02] = &RequestPlayServer{}
	packets[0x05] = &RequestServerList{}
	packets[0x07] = &RequestGGAuth{}
}

func HandlePacket(gameClient *Client, store auth.Storage, opcode uint, bytes []byte) {
	var reader = common.NewReader(bytes)

	handler := packets[opcode]

	if handler != nil {
		// TODO check here if we f*ed the performance somewhat
		ctx := client.Context{
			Client:  gameClient,
			Storage: store,
		}

		handler.HandlePacket(reader, ctx)
	} else {
		log.Printf("Failed fetching packet handler for packet %X\n", opcode)
	}
}
