package game

import (
	"l2go-concept/domain/game/packets/server"
	"l2go-concept/domain/game/storage"
	"l2go-concept/domain/network"
	"log"
)

func HandlePacket(client Client, store storage.GameStorage, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	switch opcode {
	case 0x03: // Protocol
		{
			var protocolVersion = reader.ReadUInt32()
			log.Printf("Client is with protocol version %d\n", protocolVersion)

			client.SendPacketEncoded(server.WriteKeyPacket(network.NewBuffer()))
		}
	}
}
