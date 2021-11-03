package network

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"github.com/panjf2000/gnet"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"l2go-concept/pkg/auth/crypt"
	"log"
)

type Client struct {
	blowfishKey []byte
	sessionId   uint32
	sessionKey  auth.SessionKey
	rsaKeyPair  crypt.ScrambledKeyPair
	conn        gnet.Conn
}

func (cl *Client) Options() *auth.ClientOptions {
	return &auth.ClientOptions{
		SessionId:  cl.sessionId,
		SessionKey: cl.sessionKey,
		RsaKeyPair: cl.rsaKeyPair,
		Conn:       cl.conn,
	}
}

func newClient(conn gnet.Conn) *Client {
	var blowKey = make([]byte, 16)
	var sessionId = make([]byte, 4)

	var playOk1 = make([]byte, 4)
	var playOk2 = make([]byte, 4)
	var loginOk1 = make([]byte, 4)
	var loginOK2 = make([]byte, 4)

	_, _ = rand.Read(blowKey)
	_, _ = rand.Read(sessionId)

	_, _ = rand.Read(playOk1)
	_, _ = rand.Read(playOk2)
	_, _ = rand.Read(loginOk1)
	_, _ = rand.Read(loginOK2)

	keyPair := crypt.CreateKeyPair()
	return &Client{
		blowfishKey: blowKey,
		rsaKeyPair:  keyPair,
		conn:        conn,
		sessionId:   binary.LittleEndian.Uint32(sessionId),
		sessionKey: auth.SessionKey{
			PlayOk1:  binary.LittleEndian.Uint32(playOk1),
			PlayOk2:  binary.LittleEndian.Uint32(playOk2),
			LoginOk1: binary.LittleEndian.Uint32(loginOk1),
			LoginOk2: binary.LittleEndian.Uint32(loginOK2),
		},
	}
}

func (cl *Client) SendPacketEncoded(buffer *network.Buffer) error {
	return cl.SendPacket(buffer, true, true)
}

func (cl *Client) SendPacket(inputBuffer *network.Buffer, doChecksum, doBlowfish bool) error {
	data := inputBuffer.Bytes()

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

		// Finally, do the checksum
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
	outputBuffer := network.NewBuffer()
	outputBuffer.WriteUInt16(length)
	outputBuffer.Write(data)

	bytes := outputBuffer.Bytes()

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
