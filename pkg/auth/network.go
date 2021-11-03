package auth

import (
	"github.com/panjf2000/gnet"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth/crypt"
)

type SessionKey struct {
	PlayOk1  uint32
	PlayOk2  uint32
	LoginOk1 uint32
	LoginOk2 uint32
}

type ClientOptions struct {
	SessionId  uint32
	SessionKey SessionKey
	RsaKeyPair crypt.ScrambledKeyPair
	Conn       gnet.Conn
}

type Client interface {
	SendPacketEncoded(buffer *network.Buffer) error
	SendPacket(buffer *network.Buffer, doChecksum, doBlowfish bool) error
	Receive(frame []byte) (opcode byte, data []byte, e error)
	Options() *ClientOptions
}

type Context struct {
	Client  Client
	Storage Storage
}

type ClientPacket interface {
	HandlePacket(reader *network.Reader, context Context)
}
