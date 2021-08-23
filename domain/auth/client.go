package auth

import (
	"crypto/rand"
	"github.com/panjf2000/gnet"
	"l2go-concept/domain/auth/crypt"
)

type Client struct {
	blowfishKey []byte
	rsaKeyPair  crypt.ScrambledKeyPair
	conn        gnet.Conn
}

func newClient(conn gnet.Conn) Client {
	var blowKey = make([]byte, 16)
	_, _ = rand.Read(blowKey)

	keyPair := crypt.CreateKeyPair()
	return Client{
		blowfishKey: blowKey,
		rsaKeyPair:  keyPair,
		conn:        conn,
	}
}
