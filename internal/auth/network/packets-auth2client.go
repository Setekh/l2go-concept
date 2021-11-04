package network

import (
	"l2go-concept/internal/auth/model"
	"l2go-concept/internal/common"
	"l2go-concept/pkg/auth"
	"net"
)

func clientInit(sessionId uint32, blowKey []byte, modulus []byte) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x00)

	buffer.WriteD(sessionId)                          // Session id
	buffer.WriteBytes([]byte{0x5a, 0x78, 0x00, 0x00}) // Protocol version : 785a - c4

	buffer.WriteBytes(modulus) // RSA Public

	// unk GG related?
	buffer.WriteD(0x29DD954E)
	buffer.WriteD(0x77C39CFC)
	buffer.WriteD(0x97ADB620)
	buffer.WriteD(0x07BDE0F7)

	buffer.WriteBytes(blowKey)
	buffer.WriteC(0x00) // lol

	return buffer
}

func GGAuthResponse(response uint32) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0x0B)
	buffer.WriteD(response)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)
	buffer.WriteD(0x00)

	return buffer
}

func WriteServerList(lastServer uint8, servers []*model.GameServer) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0x04)
	buffer.WriteC(byte(len(servers)))
	buffer.WriteC(lastServer)

	for _, gameServer := range servers {
		buffer.WriteC(gameServer.ServerId)

		var ip = net.ParseIP(gameServer.Ip).To4()

		buffer.WriteC(ip[0])
		buffer.WriteC(ip[1])
		buffer.WriteC(ip[2])
		buffer.WriteC(ip[3])

		buffer.WriteD(gameServer.Port)
		buffer.WriteC(0x00) // Age limit
		buffer.WriteC(0x01) // Pvp server?

		buffer.WriteH(uint16(gameServer.CurrentPlayers))
		buffer.WriteH(uint16(gameServer.MaxPlayers))

		if gameServer.IsUp {
			buffer.WriteC(0x01)
		} else {
			buffer.WriteC(0x00)
		}

		buffer.WriteD(gameServer.ServerType)
		if gameServer.ServerBrackets {
			buffer.WriteC(0x01)
		} else {
			buffer.WriteC(0x00)
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

	buffer.WriteC(0x06)
	buffer.WriteD(reason)
	return buffer
}

func LoginOk(key auth.SessionKey) *common.Buffer {
	buffer := common.NewBuffer()

	buffer.WriteC(0x03)

	buffer.WriteD(key.LoginOk1)
	buffer.WriteD(key.LoginOk2)

	buffer.WriteD(0x00)
	buffer.WriteD(0x00)

	buffer.WriteD(0x000003ea) // billing type: 1002 Free, x200 paid time, x500 flat rate pre paid, others subscription
	buffer.WriteD(0x00)       // paid time
	buffer.WriteD(0x00)
	//buffer.WriteD(0x02) - mobius??!

	buffer.WriteD(0x00)                 // warning mask
	buffer.WriteBytes(make([]byte, 16)) // forbidden servers
	//buffer.WriteD(0x00) - l2jorg
	return buffer
}

func PlayOk(serverId byte, key auth.SessionKey) *common.Buffer {
	buffer := common.NewBuffer()
	buffer.WriteC(0x07)
	buffer.WriteD(key.PlayOk1)
	buffer.WriteD(key.PlayOk2)
	buffer.WriteC(serverId) // other packs

	return buffer
}
