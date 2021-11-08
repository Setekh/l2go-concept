package client

import (
	"github.com/panjf2000/gnet"
	"l2go-concept/internal/common"
	"l2go-concept/pkg/auth"
	"l2go-concept/pkg/auth/client/crypt"
)

type SessionKey struct {
	PlayOk1  uint32
	PlayOk2  uint32
	LoginOk1 uint32
	LoginOk2 uint32
}

type Properties struct {
	SessionId  uint32
	SessionKey SessionKey
	RsaKeyPair crypt.ScrambledKeyPair
	Conn       gnet.Conn
}

type L2Client interface {
	SendPacketEncoded(buffer *common.Buffer) error
	SendPacket(buffer *common.Buffer, doChecksum, doBlowfish bool) error
	Receive(frame []byte) (opcode byte, data []byte, e error)
	Options() *Properties
}

type Context struct {
	Client  L2Client
	Storage auth.Storage
}

type IncomingPacket interface {
	HandlePacket(reader *common.Reader, context Context)
}
