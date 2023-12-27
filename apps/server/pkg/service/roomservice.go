package service

import (
	"context"

	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
)

type RoomService struct {
}

func NewRoomService() *RoomService {
	return &RoomService{}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *twirpv1.CreateRoomRequest) (*twirpv1.CreateRoomResponse, error) {
	return &twirpv1.CreateRoomResponse{SessionId: "session"}, nil
}
