package auth

import (
	"encoding/hex"
	"errors"
	"github.com/panjf2000/gnet"
	"l2go-concept/domain/auth/crypt"
	"l2go-concept/domain/auth/packets/server"
	"l2go-concept/domain/packets"
	"log"
)

var clients []Client

type clientServer struct {
	*gnet.EventServer
}

func Receive(frame []byte) (opcode byte, data []byte, e error) {
	// Read the first two bytes to define the packet size
	header := make([]byte, 2)
	//n, err := c.Socket.Read(header)
	//
	//if n != 2 || err != nil {
	//	return 0x00, nil, errors.New("An error occured while reading the packet header.")
	//}
	copy(header, frame[:2])

	// Calculate the packet size
	size := 0
	size = size + int(header[0])
	size = size + int(header[1])*256

	// Allocate the appropriate size for our data (size - 2 bytes used for the length
	data = make([]byte, size-2)
	copy(data, frame[2:])

	// Read the encrypted part of the packet
	//n, err = c.Socket.Read(data)
	//
	//if n != size-2 || err != nil {
	//	return 0x00, nil, errors.New("An error occured while reading the packet data.")
	//}

	// Print the raw packet
	log.Printf("Raw packet: \nheader:%s\n%s\n", hex.Dump(header), hex.Dump(data))

	// Decrypt the packet data using the blowfish key
	data, err := crypt.BlowfishDecrypt(data)

	if err != nil {
		return 0x00, nil, errors.New("An error occured while decrypting the packet data.")
	}

	// Verify our checksum...
	if check := crypt.Checksum(data); check {
		log.Printf("Decrypted packet content :\n%s", hex.Dump(data))
		log.Println("Packet checksum ok")
	} else {
		return 0x00, nil, errors.New("The packet checksum doesn't look right...")
	}

	// Extract the op code
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}

func (es *clientServer) React(frame []byte, c gnet.Conn) (out []byte, action gnet.Action) { // any action here is blocking
	var hexDump = hex.Dump(frame)
	log.Printf("React\n%s", hexDump)

	code, bytes, err := Receive(frame)
	if err != nil {
		log.Println("Failed decoding packet", err)
		recover()
	}

	log.Printf("Recieved code %d with decoded %s\n", code, hex.Dump(bytes))

	return
}

func SendPacket(c gnet.Conn, data []byte, doChecksum, doBlowfish bool) error {
	if doChecksum {
		// Add 4 empty bytes for the checksum new( new(
		data = append(data, []byte{0x00, 0x00, 0x00, 0x00}...)

		// Add blowfish padding
		missing := len(data) % 8

		if missing != 0 {
			for i := missing; i < 8; i++ {
				data = append(data, byte(0x00))
			}
		}

		// Finally do the checksum
		crypt.Checksum(data)
	}

	if doBlowfish {
		encrypted, err := crypt.BlowfishEncrypt(data)

		if err != nil {
			return err
		}

		data = encrypted
	}

	// Calculate the packet length
	length := uint16(len(data) + 2)

	// Put everything together
	buffer := packets.NewBuffer()
	buffer.WriteUInt16(length)
	buffer.Write(data)

	bytes := buffer.Bytes()
	//log.Printf("Sending packet bytes\n%s", hex.Dump(bytes))

	err := c.AsyncWrite(bytes)
	if err != nil {
		return err
	}

	if err != nil {
		return errors.New("the packet couldn't be sent")
	}

	return nil
}

func (es *clientServer) OnOpened(c gnet.Conn) (out []byte, action gnet.Action) {
	client := newClient(c)

	clients = append(clients, client)
	log.Println("New client!, total:", len(clients))

	buffer := packets.NewBuffer()
	server.CreateInitPacket(client.blowfishKey, client.rsaKeyPair.ScrambledModulus, buffer)
	bytes := buffer.Bytes()

	err := SendPacket(c, bytes, false, false)
	if err != nil {
		panic(err)
	}

	log.Printf("Got connection %s", c.RemoteAddr())

	//var hexDump = hex.Dump(bytes)
	//log.Printf("hex\n%s", hexDump)
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
	log.Println("Clients left", len(clients))
	return
}

func StartClientServer() {
	var server = &clientServer{}
	log.Fatalf("Server failed to start %s", gnet.Serve(server, "tcp://:2106"))
}
