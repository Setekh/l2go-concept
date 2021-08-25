package recieved

import (
	"l2go-concept/domain/network"
)

type GGAuthRequest struct {
	SessionId uint32
}

func OnGGAuth(buff *network.Reader) *GGAuthRequest {
	var sessionId = buff.ReadUInt32()

	return &GGAuthRequest{
		SessionId: sessionId,
	}
}
