package network

import (
	"encoding/hex"
	"fmt"
	auth "l2go-concept/internal/auth"
	authDef "l2go-concept/pkg/auth"
	"log"

	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
)

var clients []*Client

type clientServer struct {
	*gnet.EventServer
	pool    *goroutine.Pool
	storage authDef.Storage
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if c.Context() == nil {
		log.Printf("No recieved for this connection %s, dropping.\n", c.RemoteAddr())
		return nil, gnet.Close
	}

	client := c.Context().(*Client)

	// Handle this in another goroutine
	_ = es.pool.Submit(func() {
		onFrameDecoded(frame, client, es.storage)
	})

	return nil, gnet.None
}

func onFrameDecoded(frame []byte, client *Client, storage authDef.Storage) {
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	code, bytes, err := client.Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))

	HandlePacket(client, storage, uint(code), bytes)
}

func (es *clientServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	client := newClient(c)
	c.SetContext(client)

	clients = append(clients, client)
	log.Println("New received!, total:", len(clients))

	err := client.SendPacket(clientInit(client.sessionId, client.blowfishKey, client.rsaKeyPair.ScrambledModulus), false, false)
	if err != nil {
		log.Printf("Failed sending packet to remote connection! %s\n", c.RemoteAddr())
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

func StartClientServer(store authDef.Storage) {
	p := goroutine.Default()
	defer p.Release()

	var clientServer = &clientServer{
		pool:    p,
		storage: store,
	}

	serverConfig := auth.Config.Server
	hostname := serverConfig.Hostname
	port := serverConfig.Port

	log.Printf("Server started on address: %s:%d\n", hostname, port)
	log.Fatalf("Server failed to start %s", gnet.Serve(clientServer, fmt.Sprintf("tcp://%s:%d", hostname, port), gnet.WithReusePort(true)))
}
