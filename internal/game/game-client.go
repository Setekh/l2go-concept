package game

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
	"l2go-concept/internal/game/crypt"
	"l2go-concept/internal/network"
	"log"
)

type Client struct {
	sessionId    uint64
	cryptEnabled bool
	conn         gnet.Conn
}

func newClient(conn gnet.Conn) *Client {
	println("New client!")
	return &Client{
		conn: conn,
	}
}

func (cl *Client) SendPacket(srcBuff *network.Buffer) error {
	data := srcBuff.Bytes()

	if cl.cryptEnabled {
		log.Printf("Encoding packet %X", data[0])
		crypt.Encrypt(data)
	}

	// Calculate the packet length
	length := uint16(len(data) + 2)

	// Put everything together
	buffer := network.NewBuffer()
	buffer.WriteUInt16(length)
	buffer.Write(data)

	bytes := buffer.Bytes()
	//log.Printf("Sending packet bytes\n%s", hex.Dump(bytes))

	err := cl.conn.AsyncWrite(bytes)
	if err != nil {
		return err
	}

	if err != nil {
		return errors.New("the packet couldn't be sent")
	}

	return nil
}

func (cl *Client) Receive(frame []byte) (opcode byte, data []byte, e error) {
	header := frame[:2] // TODO this should be handled by the field length decoder

	// Calculate the packet size
	size := 0
	size = size + int(header[0])
	size = size + int(header[1])*256

	// Allocate the appropriate size for our data (size - 2 bytes used for the length
	data = frame[2:]

	// Print the raw packet
	log.Printf("Raw packet: \nheader:%s\n%s\n", hex.Dump(header), hex.Dump(data))

	if cl.cryptEnabled {
		// Decrypt the packet data using the xor key
		crypt.Decrypt(data)

		// Print the decrypted packet
		fmt.Printf("Decrypted packet content : %s\n", hex.Dump(data))
	}

	// Extract the op code
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}
