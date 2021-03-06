package network

import (
	"encoding/hex"
	"fmt"
	"github.com/panjf2000/gnet"
	"github.com/panjf2000/gnet/pool/goroutine"
	"l2go-concept/internal/game"
	"log"
)

var clients []*Client

type clientServer struct {
	*gnet.EventServer
	pool              *goroutine.Pool
	dependencyManager game.DependencyManager
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) {
	if c.Context() == nil {
		log.Printf("No recieved for this connection %s, dropping.\n", c.RemoteAddr())
		return nil, gnet.Close
	}

	client := c.Context().(*Client)

	// Handle this in another goroutine
	_ = es.pool.Submit(func() {
		onFrameDecoded(frame, client, es.dependencyManager)
	})

	return nil, gnet.None
}

func onFrameDecoded(frame []byte, client *Client, dm game.DependencyManager) {
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	code, bytes, err := client.Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))

	HandlePacket(client, dm, uint(code), bytes)
}

func (es *clientServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	client := newClient(c)
	c.SetContext(client)

	clients = append(clients, client)
	log.Println("New received!, total:", len(clients))

	log.Printf("Got connection %s", c.RemoteAddr())
	return
}

func (es *clientServer) OnClosed(conn gnet.Conn, err error) (action gnet.Action) {
	if conn.Context() == nil {
		return
	}

	client := conn.Context().(*Client)

	var index = -1
	for i, c := range clients {
		if c.sessionId == client.sessionId {
			index = i
		}
	}

	if index != -1 {
		clients = append(clients[:index], clients[index+1:]...)
	}

	log.Println("Clients left", len(clients), "error:", err)
	return
}

func StartClientServer(dependencyManager game.DependencyManager) {
	p := goroutine.Default()
	defer p.Release()

	var clientServer = &clientServer{
		pool:              p,
		dependencyManager: dependencyManager,
	}

	serverConfig := game.Config.Server
	hostname := serverConfig.Hostname
	port := serverConfig.Port

	log.Printf("Server started on address: %s:%d\n", hostname, port)
	log.Fatalf("Server failed to start %s", gnet.Serve(clientServer, fmt.Sprintf("tcp://%s:%d", hostname, port), gnet.WithReusePort(true)))
}
