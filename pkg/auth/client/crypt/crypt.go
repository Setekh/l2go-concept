package crypt

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"l2go-concept/pkg/auth/client/crypt/blowfish"
	"log"
)

type ScrambledKeyPair struct {
	PrivateKey       *rsa.PrivateKey
	PublicKey        rsa.PublicKey
	ScrambledModulus []byte
}

var cipher *blowfish.Cipher

var blowKey = []byte("_;5.]94-31==-%xT!^[$\000")

func init() {
	cip, err := blowfish.NewCipher(blowKey) // C4 client

	if err != nil {
		log.Fatalln("Failed creating cipher!", err)
	}

	cipher = cip
}

func CreateKeyPair() ScrambledKeyPair {
	privateKey, _ := rsa.GenerateKey(rand.Reader, 1024)

	scrambledModulus := privateKey.PublicKey.N.Bytes()
	scrambleModulus(scrambledModulus)

	return ScrambledKeyPair{
		PrivateKey:       privateKey,
		PublicKey:        privateKey.PublicKey,
		ScrambledModulus: scrambledModulus,
	}
}

func scrambleModulus(n []byte) {
	// step 1 : 0x4d-0x50 <-> 0x00-0x04
	for i := 0; i < 4; i++ {
		temp := n[0x00+i]
		n[0x00+i] = n[0x4d+i]
		n[0x4d+i] = temp
	}

	// step 2 xor first 0x40 bytes with last 0x40 bytes
	for i := 0; i < 0x40; i++ {
		n[i] = n[i] ^ n[0x40+i]
	}

	// step 3 xor bytes 0x0d-0x10 with bytes 0x34-0x38
	for i := 0; i < 4; i++ {
		n[0x0d+i] = n[0x0d+i] ^ n[0x34+i]
	}

	// step 4 xor last 0x40 bytes with first 0x40 bytes
	for i := 0; i < 0x40; i++ {
		n[0x40+i] = n[0x40+i] ^ n[i]
	}
}

func Checksum(raw []byte) bool {
	var chksum = 0
	count := len(raw) - 8
	i := 0

	for i = 0; i < count; i += 4 {
		var ecx = int(raw[i])
		ecx |= int(raw[i+1]) << 8
		ecx |= int(raw[i+2]) << 0x10
		ecx |= int(raw[i+3]) << 0x18
		chksum ^= ecx
	}

	var ecx = int(raw[i])
	ecx |= int(raw[i+1]) << 8
	ecx |= int(raw[i+2]) << 0x10
	ecx |= int(raw[i+3]) << 0x18

	raw[i] = byte(chksum)
	raw[i+1] = byte(chksum >> 0x08)
	raw[i+2] = byte(chksum >> 0x10)
	raw[i+3] = byte(chksum >> 0x18)

	return ecx == chksum
}

func BlowfishDecrypt(encrypted []byte) ([]byte, error) {
	// Check if the encrypted data is a multiple of our block size
	if len(encrypted)%8 != 0 {
		return nil, errors.New("the encrypted data is not a multiple of the block size")
	}

	count := len(encrypted) / 8

	decrypted := make([]byte, len(encrypted))

	for i := 0; i < count; i++ {
		cipher.Decrypt(decrypted[i*8:], encrypted[i*8:])
	}

	return decrypted, nil
}

func BlowfishEncrypt(decrypted []byte) ([]byte, error) {
	// Check if the decrypted data is a multiple of our block size
	if len(decrypted)%8 != 0 {
		return nil, errors.New("the decrypted data is not a multiple of the block size")
	}

	count := len(decrypted) / 8

	encrypted := make([]byte, len(decrypted))

	for i := 0; i < count; i++ {
		cipher.Encrypt(encrypted[i*8:], decrypted[i*8:])
	}

	return encrypted, nil
}
