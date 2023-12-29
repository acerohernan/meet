package service

import (
	"context"

	"github.com/acerohernan/meet/core"
	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
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

	err = s.router.SendNodeMessage(node.Id, &core.NodeMessage{
		Message: &core.NodeMessage_CreateRoom{
			CreateRoom: &core.CreateRoom{
				RoomId: "roomID",
			},
		},
	})

	if err != nil {
		return nil, err
	}

	//  confirm execution

	return &twirpv1.CreateRoomResponse{AccessToken: "xd"}, nil
}
