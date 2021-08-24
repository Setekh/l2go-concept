package auth

import (
	"encoding/hex"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/packets"
	"log"
)

var clients []Client

type clientServer struct {
	*gnet.EventServer
	pool *goroutine.Pool
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if c.Context() == nil {
		log.Printf("No client for this connection %s, dropping.\n", c.RemoteAddr())
		_ = c.Close()
		return nil, gnet.Close
	}

	client := c.Context().(Client)

	// Handle this in another goroutine
	_ = es.pool.Submit(func() {
		handlePacket(frame, client)
	})

	return nil, gnet.None
}

func handlePacket(frame []byte, client Client) {
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	code, bytes, err := client.Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
		recover()
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))
}

func (es *clientServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	client := newClient(c)
	c.SetContext(client)

	clients = append(clients, client)
	log.Println("New client!, total:", len(clients))

	buffer := packets.NewBuffer()
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

func StartClientServer() {
	p := goroutine.Default()
	defer p.Release()

	var clientServer = &clientServer{
		pool: p,
	}
	log.Fatalf("Server failed to start %s", gnet.Serve(clientServer, "tcp://:2106", gnet.WithReusePort(true)))
}
