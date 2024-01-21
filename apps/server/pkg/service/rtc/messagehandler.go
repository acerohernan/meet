package rtc

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

func (m *rtcManager) handleNodeMessage(msg *core.NodeMessage) error {
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

	case *core.NodeMessage_GuestJoinRequest:
		req := msg.GetGuestJoinRequest()
		room, err := m.GetRoom(m.ctx, req.RoomId)
		if err != nil {
			return err
		}

		room.AddGuest(req.Guest)
		room.SendNewGuestRequest(req.Guest)

	case *core.NodeMessage_GuestRequestCancelled:
		req := msg.GetGuestRequestCancelled()
		room, err := m.GetRoom(m.ctx, req.RoomId)
		if err != nil {
			return err
		}
		room.DeleteGuest(req.GuestId)
		room.SendGuestRequestCancelled(req.GuestId)
	}

	return nil
}
