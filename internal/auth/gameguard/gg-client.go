package gameguard

import "l2go-concept/internal/network"

type GGAuthRequest struct {
	SessionId uint32
}

func RequestGGAuth(buff *network.Reader) *GGAuthRequest {
	var sessionId = buff.ReadUInt32()

	return &GGAuthRequest{
		SessionId: sessionId,
	}
}
