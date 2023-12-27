package service

import (
	"context"

	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/service/auth"
)

type RoomService struct {
	authSvc *auth.AuthService
}

func NewRoomService(authSvc *auth.AuthService) *RoomService {
	return &RoomService{
		authSvc: authSvc,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *twirpv1.CreateRoomRequest) (*twirpv1.CreateRoomResponse, error) {
	return &twirpv1.CreateRoomResponse{SessionId: "session"}, nil
}
