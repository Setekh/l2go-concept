package network

import (
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/auth/storage"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
	"math/big"
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
	var publicIp = network.GetPublicIp()

	gs := &model.GameServer{
		ServerId:   0x02,
		Ip:         publicIp,
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

	reader := network.NewReader(decodedBytes)
	reader.Seek(3, 0)
	var accountName = reader.ReadString() //strings.TrimSpace(string(decodedBytes[3:17]))
	reader.Seek(17, 0)
	var password = reader.ReadString() //strings.TrimSpace(string(decodedBytes[17:]))

	result := store.VerifyAccount(accountName, password)
	log.Printf("User %s is trying to connect with password %s %d result", accountName, password, result)

	if result == storage.AccountNotFound || result == storage.Ok {
		store.CreateAccount(accountName, password)
		log.Printf("Created account for user %s", accountName)
		client.SendPacketEncoded(LoginOk(properties.SessionKey))
	} else if result == storage.InvalidPassword {
		client.SendPacketEncoded(LoginFail(AccountPasswordWrong))
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
	client.SendPacketEncoded(PlayOk(serverId, options.SessionKey))
}
