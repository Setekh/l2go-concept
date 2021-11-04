package network

import (
	bytes2 "bytes"
	"encoding/hex"
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/auth/storage"
	"l2go-concept/internal/common"
	"l2go-concept/pkg/auth"
	"log"
	"math/big"
	"os"
)

type RequestGGAuth struct{}

func (p *RequestGGAuth) HandlePacket(buff *common.Reader, ctx auth.Context) {
	var sessionId = buff.ReadD()

	options := ctx.Client.Options()
	log.Printf("GG trys to authenticate! sessionId %d == %d\n", sessionId, options.SessionId)

	if sessionId != options.SessionId {
		options.Conn.Close()
		return
	}

	ctx.Client.SendPacketEncoded(GGAuthResponse(sessionId))
}

type RequestServerList struct{}

func (p *RequestServerList) HandlePacket(_ *common.Reader, ctx auth.Context) {
	serverIp := os.Getenv("game.server.address")

	var publicIp string

	if serverIp != "" {
		publicIp = serverIp
	} else {
		publicIp = common.GetPublicIp()
	}

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

func (p *RequestAuth) HandlePacket(buff *common.Reader, ctx auth.Context) {
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

	println(hex.Dump(decodedBytes[3:17]))
	println(hex.Dump(decodedBytes[17:]))

	readTillStop := func(bytes []byte) string {
		indexByte := bytes2.IndexByte(bytes, 0x00)
		return string(bytes[:indexByte])
	}

	var accountName = readTillStop(decodedBytes[3:17])
	var password = readTillStop(decodedBytes[17:])

	result := store.VerifyAccount(accountName, password)
	log.Printf("User %s is trying to connect with password %s %d result", accountName, password, result)

	if result == storage.InvalidPassword {
		client.SendPacketEncoded(LoginFail(AccountPasswordWrong))
		return
	}

	if result == storage.AccountNotFound {
		store.CreateAccount(accountName, password)
		log.Printf("Created account for user %s", accountName)
	}

	client.SendPacketEncoded(LoginOk(properties.SessionKey))
}

type RequestPlayServer struct{}

func (p *RequestPlayServer) HandlePacket(buff *common.Reader, ctx auth.Context) {
	client := ctx.Client
	options := client.Options()

	loginOK1 := buff.ReadD()
	loginOK2 := buff.ReadD()
	serverId, _ := buff.ReadByte()

	println("Received login oks", loginOK1, loginOK2, "wants to join", serverId)
	client.SendPacketEncoded(PlayOk(serverId, options.SessionKey))
}
