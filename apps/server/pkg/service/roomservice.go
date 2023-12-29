package service

import (
	"context"

	"github.com/acerohernan/meet/core"
	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/rtc"
	"github.com/acerohernan/meet/pkg/utils"
)

type RoomService struct {
	router  *router.Router
	manager *rtc.RoomManager
}

func NewRoomService(router *router.Router, manager *rtc.RoomManager) *RoomService {
	return &RoomService{
		router:  router,
		manager: manager,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *twirpv1.CreateRoomRequest) (*twirpv1.CreateRoomResponse, error) {
	// select a node to host the room by cpu usage
	node, err := s.router.GetAvaliableNode()

	// TODO: send a twirp http error
	if err != nil {
		return nil, err
	}

	roomId := utils.NewId(utils.RoomPrefix)

	// send node message
	err = s.router.SendNodeMessage(node.Id, &core.NodeMessage{
		Message: &core.NodeMessage_CreateRoom{
			CreateRoom: &core.CreateRoom{
				RoomId: roomId,
			},
		},
	})

	if err != nil {
		return nil, err
	}

	// TODO: confirm execution

	return &twirpv1.CreateRoomResponse{AccessToken: "token", RoomId: roomId}, nil
}
