package server

import (
	"l2go-concept/domain/auth/model"
	"l2go-concept/domain/network"
	"net"
)

func WriteServerList(lastServer uint8, servers []model.GameServer) *network.Buffer {
	buffer := network.NewBuffer()

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
