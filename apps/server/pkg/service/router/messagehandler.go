package router

import (
	"github.com/acerohernan/meet/core"
	"github.com/acerohernan/meet/pkg/config/logger"
)

func (r *Router) handleNodeMessage(msg *core.NodeMessage) error {
	switch msg.Message.(type) {
	case *core.NodeMessage_CreateRoom:
		logger.Infow("create room node message received!", "msg", msg)

		req := msg.GetCreateRoom()

		if err := r.manager.CreateRoom(r.ctx, req.RoomId); err != nil {
			return err
		}
	}

	return nil
}
