package login

import (
	"l2go-concept/internal/auth/packets/server"
	"l2go-concept/internal/auth/storage"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
)

type PacketHandler struct{}

func (p *PacketHandler) HandlePacket(_ uint, client auth.Client, reader *network.Reader, store auth.Storage) {
	properties := client.Properties()

	userCredentials, err := RequestAuth(properties.RsaKeyPair.PrivateKey, reader)
	if err != nil {
		properties.Conn.Close()
		return
	}

	result := store.VerifyAccount(userCredentials.AccountName, userCredentials.Password)
	log.Printf("User %s is trying to connect %d result", userCredentials.AccountName, result)

	if result == storage.AccountNotFound || result == storage.Ok {
		store.CreateAccount(userCredentials.AccountName, userCredentials.Password)
		log.Printf("Created account for user %s", userCredentials.AccountName)
		client.SendPacketEncoded(server.WriteLoginOk(properties.SessionId, properties.SessionKey))
	} else if result == storage.InvalidPassword {
		client.SendPacketEncoded(server.WriteLoginFail(server.AccountPasswordWrong))
	}
}
