package auth

import (
	"l2go-concept/domain/auth/packets/recieved"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/network"
	"log"
)

func HandlePacket(client Client, opcode uint, bytes []byte) {
	var reader = network.NewReader(bytes)

	switch opcode {
	case 0x07:
		{
			authRequest := recieved.OnGGAuth(reader)
			reply := server.GGAuthResponse(authRequest.SessionId)
			client.SendPacketEncoded(reply)
		}
	case 0x00:
		{
			userCredentials, err := recieved.OnRequestAuth(client.rsaKeyPair.PrivateKey, reader)
			if err != nil {
				client.conn.Close()
				return
			}

			log.Printf("User %s is trying to connect", userCredentials.AccountName)
		}
	}
}
