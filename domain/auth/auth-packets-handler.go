package auth

import (
	"l2go-concept/domain/auth/model"
	"l2go-concept/domain/auth/packets/recieved"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/auth/storage"
	"l2go-concept/domain/network"
	"log"
)

func HandlePacket(client Client, store storage.LoginStorage, opcode uint, bytes []byte) {
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

			result := store.VerifyAccount(userCredentials.AccountName, userCredentials.Password)
			log.Printf("User %s is trying to connect %d result", userCredentials.AccountName, result)

			if result == storage.AccountNotFound || result == storage.Ok {
				store.CreateAccount(userCredentials.AccountName, userCredentials.Password)
				log.Printf("Created account for user %s", userCredentials.AccountName)
				client.SendPacketEncoded(server.WriteLoginOk(client.sessionId))
			} else if result == storage.InvalidPassword {
				client.SendPacketEncoded(server.WriteLoginFail(server.AccountPasswordWrong))
			}
		}
	case 0x05: // Request server list
		{
			// TODO extract this in a toml?
			gs := &model.GameServer{
				ServerId:   0x01,
				Ip:         "192.168.100.69",
				Port:       7777,
				MaxPlayers: 1000,
				PvpServer:  true,
				IsUp:       true,
			}

			var gsList = []model.GameServer{*gs}
			client.SendPacketEncoded(server.WriteServerList(0x00, gsList))
		}
	case 0x02: // Request play
		{
			client.SendPacketEncoded(server.WritePlayOk(client.sessionId))
		}
	}
}
