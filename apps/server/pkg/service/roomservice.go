package service

import (
	"context"

	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/router"
)

type RoomService struct {
	router *router.Router
}

func NewRoomService(router *router.Router) *RoomService {
	return &RoomService{
		router: router,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *twirpv1.CreateRoomRequest) (*twirpv1.CreateRoomResponse, error) {
	// select a node to host the room by cpu usage
	node, err := s.router.GetAvaliableNode()

	if err != nil {
		return nil, err
	}

	logger.Infow("node found", "node", node)

	// I need to have a NODE structure that will store all tbe node stats, an I will have a go routine that each 5 seconds will write down the instance stats

	/*
		node, err := router.GetAvaliableNode()

		if err != nil && err != NotFoundAvaliableNodeErr {
			return nil, err
		}

		if err ==  NotFoundAvaliableNodeErr{
			// feedback
		}

		// comunicate via redis queue that a room wants to be created in the specific node
		err := node.SendNodeMessage(*core.Message{
			Message: *core.Message_CreateRoom{}
		})

		// wait until it is completed
		2 seconds timeout
		confirm execution each 100 ms
	*/

	return &twirpv1.CreateRoomResponse{AccessToken: "xd"}, nil
}
