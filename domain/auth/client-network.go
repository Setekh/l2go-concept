package auth

import (
	"encoding/hex"
	"github.com/panjf2000/gnet"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/packets"
	"log"
)

var clients []Client

type clientServer struct {
	*gnet.EventServer
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) { // any action here is blocking
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	client := getClient(c)

	if client == nil {
		log.Printf("No client for this connection %s, dropping.\n", c.RemoteAddr())
		return nil, gnet.Close
	}

	code, bytes, err := client.Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
		recover()
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))

	return
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

func getClient(c gnet.Conn) *Client {
	context := c.Context()

	switch client := context.(type) {
	case Client:
		return &client
	}

	return nil
}

func StartClientServer() {
	var clientServer = &clientServer{}
	log.Fatalf("Server failed to start %s", gnet.Serve(clientServer, "tcp://:2106"))
}
