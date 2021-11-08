package network

import (
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/panjf2000/gnet"
	"l2go-concept/internal/common"
	"l2go-concept/internal/game/model"
	"l2go-concept/internal/game/network/crypt"
	"l2go-concept/pkg/game/client"
	"log"
)

type Client struct {
	cryptEnabled bool
	crypt        crypt.Crypt
	conn         gnet.Conn
	sessionId    uint32
	accountName  string
	player       *model.Character
}

func newClient(conn gnet.Conn) *Client {
	println("New client!")
	return &Client{
		conn: conn,
		crypt: crypt.Crypt{
			InputKey:  crypt.GetKey(),
			OutputKey: crypt.GetKey(),
		},
		cryptEnabled: false,
	}
}

func (cl *Client) GetSessionId() uint32 {
	return cl.sessionId
}

func (cl *Client) GetAccountName() string {
	return cl.accountName
}

func (cl *Client) GetPlayer() *model.Character {
	return cl.player
}

func (cl *Client) EnableCrypt() {
	cl.cryptEnabled = true
}

func (cl *Client) Upgrade(accountName string, sessionId uint32) {
	cl.accountName = accountName
	cl.sessionId = sessionId
}

func (cl *Client) SetPlayer(character *model.Character) {
	cl.player = character
}

func (cl *Client) SendPacket(packet client.OutgoingPacket, params ...interface{}) {
	buffer := common.NewBuffer()
	packet.WritePacket(buffer, params...)
	err := cl.SendRawPacket(buffer)
	if err != nil {
		log.Printf("Failed sending packet!\n%s\n", hex.Dump(buffer.Bytes()))
	}
}

func (cl *Client) Close() {
	cl.conn.Close()
}

func (cl *Client) SendRawPacket(srcBuff *common.Buffer) error {
	data := srcBuff.Bytes()

	if cl.cryptEnabled {
		log.Printf("Encoding packet %X", data[0])
		cl.crypt.Encrypt(data)
	}

	// Calculate the packet length
	length := uint16(len(data) + 2)

	// Put everything together
	buffer := common.NewBuffer()
	buffer.WriteH(length)
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
		cl.crypt.Decrypt(data)

		// Print the decrypted packet
		fmt.Printf("Decrypted packet content : %s\n", hex.Dump(data))
	}

	// Extract the op code
	opcode = data[0]
	data = data[1:]
	e = nil
	return
}
