package clientPackets

import (
	"l2go-concept/internal/common"
	"l2go-concept/internal/game"
	"l2go-concept/pkg/game/client"
)

type RequestLogout struct{}

func (r RequestLogout) ReadPacket(client client.L2Client, _ game.DependencyManager, _ *common.Reader) {
	client.SendPacket(LogoutStatic)
}

type logout struct{}

var LogoutStatic = &logout{}

func (l *logout) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	buffer.WriteC(0x7e)
}

type RequestRestart struct{}

func (r RequestRestart) ReadPacket(client client.L2Client, dm game.DependencyManager, reader *common.Reader) {
	client.SendPacket(RestartReplaySuccessful)

	characters := dm.GetStorage().LoadAllCharacters(client.GetAccountName())
	client.SendPacket(&CharacterList{
		Characters:  characters,
		AccountName: client.GetAccountName(),
		SessionId:   client.GetSessionId(),
	})
}

type RequestRestartReplay struct {
	successful bool
}

var RestartReplaySuccessful = &RequestRestartReplay{
	successful: true,
}

func (r *RequestRestartReplay) WritePacket(buffer *common.Buffer, _ ...interface{}) {
	buffer.WriteC(0x5F)
	if r.successful {
		buffer.WriteD(0x01)
	} else {
		buffer.WriteD(0x00)
	}
}
