package login

import (
	"crypto/rsa"
	"errors"
	"l2go-concept/internal/network"
	"math/big"
	"strings"
)

type UserCredentials struct {
	AccountName string
	Password    string
}

func RequestAuth(key *rsa.PrivateKey, buff *network.Reader) (*UserCredentials, error) {
	if buff.Len() <= 128 {
		return nil, errors.New("invalid request auth packet, too small")
	}

	bytes := buff.ReadBytes(128)
	encoded := new(big.Int).SetBytes(bytes)
	decodedBytes := encoded.Exp(encoded, key.D, key.N).Bytes()

	var username = strings.TrimSpace(string(decodedBytes[3:17]))
	var password = strings.TrimSpace(string(decodedBytes[17:]))

	return &UserCredentials{
		AccountName: strings.ToLower(username),
		Password:    password,
	}, nil
}
