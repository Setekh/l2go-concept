package network

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/network/crypt"
	"l2go-concept/internal/game/network/server"
	"l2go-concept/internal/game/storage"
	"log"
)

func HandlePacket(client *Client, store storage.GameStorage, opcode uint, bytes []byte) {
	var reader = common.NewReader(bytes)

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
			RequestCharacterList(client, store, reader)
			break
		}
	case 0x09: // Logout
		{
			buffer := common.NewBuffer()
			buffer.WriteC(0x7e)
			client.SendPacket(buffer)
			//client.conn.Close()
			break
		}
	case 0x0e: // Create new Character
		{
			buffer := common.NewBuffer()
			RequestCreateCharacter(client, buffer)
			break
		}
	case 0x0b: // Request Create Character
		{
			CreateCharacter(client, store, common.NewReader(bytes))
		}
	}
}
