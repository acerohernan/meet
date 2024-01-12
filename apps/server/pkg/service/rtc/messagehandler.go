package rtc

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

func (m *rtcManager) handleNodeMessage(msg *core.NodeMessage) error {
	switch msg.Message.(type) {
	case *core.NodeMessage_CreateRoom:
		logger.Infow("creating room...", "msg", msg)
		if err := m.CreateRoom(m.ctx, msg.GetCreateRoom().RoomId); err != nil {
			return err
		}
	}

	return nil
}
