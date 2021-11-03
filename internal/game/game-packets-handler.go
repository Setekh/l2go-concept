package game

import (
	"l2go-concept/internal/game/crypt"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/packets/server"
	"l2go-concept/internal/game/storage"
	"l2go-concept/internal/network"
	"log"
)

func HandlePacket(client *Client, store storage.GameStorage, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	switch opcode {
	case 0x00: // Protocol
		{
			var protocolVersion = reader.ReadUInt32()
			log.Printf("Client is with protocol version %d\n", protocolVersion)

			client.SendPacket(server.WriteKeyPacket(crypt.GetKey()))
			client.cryptEnabled = true
			break
		}
	case 0x08: // Request auth
		{
			accountName := reader.ReadString()
			playOk2 := reader.ReadUInt32()
			playOk1 := reader.ReadUInt32()
			loginOk1 := reader.ReadUInt32()
			loginOk2 := reader.ReadUInt32()

			println(accountName, loginOk1, loginOk2, playOk1, playOk2)

			client.SendPacket(server.WriteCharacterList(make([]model.Character, 0)))
			break
		}
	case 0x09: // Logout
		{
			buffer := network.NewBuffer()
			buffer.WriteByte(0x7e)
			client.SendPacket(buffer)
			//client.conn.Close()
			break
		}
	}
}
