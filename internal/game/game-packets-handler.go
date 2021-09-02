package game

import (
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/packets/server"
	"l2go-concept/internal/game/storage"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
)

func HandlePacket(client *Client, store storage.GameStorage, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	switch opcode {
	case 0x00: // Protocol
		{
			var protocolVersion = reader.ReadUInt32()
			log.Printf("Client is with protocol version %d\n", protocolVersion)

			client.SendPacket(server.WriteKeyPacket())
			client.cryptEnabled = true
		}
	case 0x08: // Request auth
		{
			accountName := reader.ReadString()
			playOk2 := reader.ReadUInt32()
			playOk1 := reader.ReadUInt32()
			loginOk1 := reader.ReadUInt32()
			loginOk2 := reader.ReadUInt32()

			println(accountName, loginOk1, loginOk2, playOk1, playOk2)

			client.SendPacket(server.WriteSeasonKey(accountName, auth.SessionKey{
				PlayOk1:  playOk1,
				PlayOk2:  playOk2,
				LoginOk1: loginOk1,
				LoginOk2: loginOk2,
			}))

			client.SendPacket(server.WriteCharacterList(make([]model.Character, 0)))
		}
	}
}
