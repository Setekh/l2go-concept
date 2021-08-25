package auth

import (
	"encoding/hex"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/auth/storage"
	"l2go-concept/domain/network"
	"log"
)

var clients []Client

type clientServer struct {
	*gnet.EventServer
	pool    *goroutine.Pool
	storage storage.LoginStorage
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if c.Context() == nil {
		log.Printf("No recieved for this connection %s, dropping.\n", c.RemoteAddr())
		return nil, gnet.Close
	}

	client := c.Context().(Client)

	// Handle this in another goroutine
	_ = es.pool.Submit(func() {
		onFrameDecoded(frame, client, es.storage)
	})

	return nil, gnet.None
}

func onFrameDecoded(frame []byte, client Client, storage storage.LoginStorage) {
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	code, bytes, err := client.Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
		recover()
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))

	HandlePacket(client, storage, uint(code), bytes)
}

func (es *clientServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	client := newClient(c)
	c.SetContext(client)

	clients = append(clients, client)
	log.Println("New received!, total:", len(clients))

	buffer := network.NewBuffer()
	server.CreateInitPacket(client.blowfishKey, client.rsaKeyPair.ScrambledModulus, buffer)
	bytes := buffer.Bytes()

	err := client.SendPacket(bytes, false, false)
	if err != nil {
		panic(err)
	}

	log.Printf("Got connection %s", c.RemoteAddr())
	return
}

func (es *clientServer) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
	var index = -1

	for i, c := range clients {
		if c.conn.RemoteAddr() == conn.RemoteAddr() {
			index = i
		}
	}

	if index != -1 {
		clients = append(clients[:index], clients[index+1:]...)
	}

	log.Println("Clients left", len(clients), "error:", err)
	return
}

func StartClientServer(store storage.LoginStorage) {
	p := goroutine.Default()
	defer p.Release()

	var clientServer = &clientServer{
		pool:    p,
		storage: store,
	}
	log.Fatalf("Server failed to start %s", gnet.Serve(clientServer, "tcp://:2106", gnet.WithReusePort(true)))
}
