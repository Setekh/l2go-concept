package server

import (
	"encoding/hex"
	"l2go-concept/internal/network"
	"l2go-concept/pkg/auth"
	"log"
	"unicode/utf16"
)

func utf16leMd5(s string) []byte {
	codes := utf16.Encode([]rune(s))
	b := make([]byte, len(codes)*2)
	for i, r := range codes {
		b[i*2] = byte(r)
		b[i*2+1] = byte(r >> 8)
	}
	return b
}

func WriteSeasonKey(account string, key auth.SessionKey) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x05)

	accountBytes := []rune(account)

	utf16.Encode(accountBytes)

	//TODO move this into a method
	for _, accountRune := range accountBytes {
		buffer.WriteRune(accountRune)
	}

	buffer.WriteByte(0x00)
	buffer.WriteByte(0x00)

	buffer.WriteUInt32(uint32(key.LoginOk1))
	buffer.WriteUInt32(uint32(key.LoginOk2))
	buffer.WriteUInt32(uint32(key.PlayOk1))
	buffer.WriteUInt32(uint32(key.PlayOk2))

	log.Printf("Sending auth reply!\n%s\n", hex.Dump(buffer.Bytes()))
	return buffer
}

func WriteSeasonKey2(account string, key uint64) *network.Buffer {
	buffer := network.NewBuffer()

	buffer.WriteByte(0x05)

	accountBytes := []rune(account)
	utf16.Encode(accountBytes)

	//TODO move this into a method
	for _, accountRune := range accountBytes {
		buffer.WriteRune(accountRune)
	}
	buffer.WriteByte(0x00)
	buffer.WriteByte(0x00)

	buffer.WriteUInt64(key)
	return buffer
}
