package network

import (
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/auth/packets/server"
	"l2go-concept/internal/auth/storage"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
	"math/big"
	"strings"
)

type RequestGGAuth struct{}

func (p *RequestGGAuth) HandlePacket(buff *network.Reader, ctx auth.Context) {
	var sessionId = buff.ReadUInt32()

	options := ctx.Client.Options()
	log.Printf("GG trys to authenticate! sessionId %d == %d\n", sessionId, options.SessionId)

	if sessionId != options.SessionId {
		options.Conn.Close()
		return
	}

	ctx.Client.SendPacketEncoded(GGAuthResponse(sessionId))
}

type RequestServerList struct{}

func (p *RequestServerList) HandlePacket(_ *network.Reader, ctx auth.Context) {
	gs := &model.GameServer{
		ServerId:   0x02,
		Ip:         "127.0.0.1",
		Port:       7777,
		MaxPlayers: 1000,
		PvpServer:  true,
		IsUp:       true,
	}

	// This will be fetched from the db most probably
	list := []*model.GameServer{gs}

	ctx.Client.SendPacketEncoded(WriteServerList(0x00, list))
}

type RequestAuth struct{}

func (p *RequestAuth) HandlePacket(buff *network.Reader, ctx auth.Context) {
	client := ctx.Client
	store := ctx.Storage
	properties := client.Options()

	if buff.Len() <= 128 {
		properties.Conn.Close()
		return
	}

	bytes := buff.ReadBytes(128)
	encoded := new(big.Int).SetBytes(bytes)

	key := properties.RsaKeyPair.PrivateKey
	decodedBytes := encoded.Exp(encoded, key.D, key.N).Bytes()

	var accountName = strings.TrimSpace(string(decodedBytes[3:17]))
	var password = strings.TrimSpace(string(decodedBytes[17:]))

	result := store.VerifyAccount(accountName, password)
	log.Printf("User %s is trying to connect with password %s %d result", accountName, password, result)

	if result == storage.AccountNotFound || result == storage.Ok {
		store.CreateAccount(accountName, password)
		log.Printf("Created account for user %s", accountName)
		client.SendPacketEncoded(server.WriteLoginOk(properties.SessionKey))
	} else if result == storage.InvalidPassword {
		client.SendPacketEncoded(server.WriteLoginFail(server.AccountPasswordWrong))
	}
}

type RequestPlayServer struct{}

func (p *RequestPlayServer) HandlePacket(buff *network.Reader, ctx auth.Context) {
	client := ctx.Client
	options := client.Options()

	loginOK1 := buff.ReadUInt32()
	loginOK2 := buff.ReadUInt32()
	serverId, _ := buff.ReadByte()

	println("Received login oks", loginOK1, loginOK2, "wants to join", serverId)
	println("Yoo", options.SessionId, "heh", options.SessionKey.LoginOk1, options.SessionKey.LoginOk2, options.SessionKey.PlayOk1, options.SessionKey.PlayOk2)
	client.SendPacketEncoded(server.WritePlayOk(serverId, options.SessionKey))
}
