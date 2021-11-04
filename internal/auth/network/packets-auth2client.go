package network

import (
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/common"
	"l2go-concept/pkg/auth"
	"net"
)

func clientInit(sessionId uint32, blowKey []byte, modulus []byte) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteByte(0x00)

	buffer.WriteUInt32(sessionId)                // Session id
	buffer.Write([]byte{0x5a, 0x78, 0x00, 0x00}) // Protocol version : 785a - c4

	buffer.Write(modulus) // RSA Public

	// unk GG related?
	buffer.WriteUInt32(0x29DD954E)
	buffer.WriteUInt32(0x77C39CFC)
	buffer.WriteUInt32(0x97ADB620)
	buffer.WriteUInt32(0x07BDE0F7)

	buffer.Write(blowKey)
	buffer.WriteByte(0x00) // lol

	return buffer
}

func GGAuthResponse(response uint32) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteByte(0x0B)
	buffer.WriteUInt32(response)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)

	return buffer
}

func WriteServerList(lastServer uint8, servers []*model.GameServer) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteByte(0x04)
	buffer.WriteByte(byte(len(servers)))
	buffer.WriteByte(lastServer)

	for _, gameServer := range servers {
		buffer.WriteByte(gameServer.ServerId)

		var ip = net.ParseIP(gameServer.Ip).To4()

		buffer.WriteByte(ip[0])
		buffer.WriteByte(ip[1])
		buffer.WriteByte(ip[2])
		buffer.WriteByte(ip[3])

		buffer.WriteUInt32(gameServer.Port)
		buffer.WriteByte(0x00) // Age limit
		buffer.WriteByte(0x01) // Pvp server?

		buffer.WriteUInt16(uint16(gameServer.CurrentPlayers))
		buffer.WriteUInt16(uint16(gameServer.MaxPlayers))

		if gameServer.IsUp {
			buffer.WriteByte(0x01)
		} else {
			buffer.WriteByte(0x00)
		}

		buffer.WriteUInt32(gameServer.ServerType)
		if gameServer.ServerBrackets {
			buffer.WriteByte(0x01)
		} else {
			buffer.WriteByte(0x00)
		}
	}

	return buffer
}

const (
	_NONE = iota
	SystemError
	AccountPasswordWrong
	AccountOrPasswordWrong
	AccessFailed
	AccountInUse = 0x07
)

func LoginFail(reason uint32) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteByte(0x06)
	buffer.WriteUInt32(reason)
	return buffer
}

func LoginOk(key auth.SessionKey) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteByte(0x03)

	buffer.WriteUInt32(key.LoginOk1)
	buffer.WriteUInt32(key.LoginOk2)

	buffer.WriteUInt32(0x00)
	buffer.WriteUInt32(0x00)

	buffer.WriteUInt32(0x000003ea) // billing type: 1002 Free, x200 paid time, x500 flat rate pre paid, others subscription
	buffer.WriteUInt32(0x00)       // paid time
	buffer.WriteUInt32(0x00)
	//buffer.WriteUInt32(0x02) - mobius??!

	buffer.WriteUInt32(0x00)       // warning mask
	buffer.Write(make([]byte, 16)) // forbidden servers
	//buffer.WriteUInt32(0x00) - l2jorg
	return buffer
}

func PlayOk(serverId byte, key auth.SessionKey) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteByte(0x07)
	buffer.WriteUInt32(key.PlayOk1)
	buffer.WriteUInt32(key.PlayOk2)
	buffer.WriteByte(serverId) // other packs

	return buffer
}
