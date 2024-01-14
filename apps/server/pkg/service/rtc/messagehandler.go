package rtc

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

func (m *rtcManager) handleNodeMessage(msg *core.NodeMessage) error {
	logger.Infow("node message received", "msg", msg)
	switch msg.Message.(type) {
	case *core.NodeMessage_CreateRoom:
		if err := m.CreateRoom(m.ctx, msg.GetCreateRoom().RoomId); err != nil {
			return err
		}

	case *core.NodeMessage_SignalRequest:
		req := msg.GetSignalRequest()
		room, err := m.GetRoom(m.ctx, req.RoomId)
		if err != nil {
			return err
		}
		p := room.GetParticipant(req.ParticipantId)

		if err := m.handleSignalRequest(room, p, req); err != nil {
			logger.Errorw("error at handling signal request", err)
		}

	case *core.NodeMessage_SignalResponse:
		res := msg.GetSignalResponse()
		m.ReceiveParticipantResponse(res.ParticipantId, res)
	}

	return nil
}
