package service

import (
	"context"
	"errors"
	"time"

	"github.com/acerohernan/meet/core"
	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/storage"
	"github.com/acerohernan/meet/pkg/utils"
	"github.com/twitchtv/twirp"
)

var (
	APITimeoutDuration       = time.Second * 2
	ConfirmExecutionInterval = time.Millisecond * 200

	APITimeoutError = errors.New("api timeout")
)

type RoomService struct {
	router  *router.Router
	store   storage.ObjectStore
	authSvc *auth.AuthService
}

func NewRoomService(router *router.Router, store storage.ObjectStore, authSvc *auth.AuthService) *RoomService {
	return &RoomService{
		router:  router,
		store:   store,
		authSvc: authSvc,
	}
}

func (s *RoomService) CreateRoom(ctx context.Context, req *twirpv1.CreateRoomRequest) (*twirpv1.CreateRoomResponse, error) {
	// select a node to host the room by cpu usage
	node, err := s.router.GetAvaliableNode()

	if err != nil {
		return nil, twirp.NewError(twirp.Unavailable, "could find an avaliable node to host the room")
	}

	if req.RoomId == "" {
		req.RoomId = utils.NewId(utils.RoomPrefix)
	}

	// send node message
	err = s.router.SendNodeMessage(node.Id, &core.NodeMessage{
		Message: &core.NodeMessage_CreateRoom{
			CreateRoom: &core.CreateRoom{
				RoomId: req.RoomId,
			},
		},
	})

	if err != nil {
		return nil, twirp.NewError(twirp.Internal, "could not communicate with room to host")
	}

	err = confirmExecution(func() error {
		_, err := s.store.LoadRoom(ctx, req.RoomId)
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		if err == APITimeoutError {
			return nil, twirp.NewError(twirp.Internal, "api call timeout")
		}

		return nil, twirp.NewError(twirp.Internal, "couldnt confirm message execution")
	}

	// create access token for the new created room
	grants := &auth.Grants{
		ID:        utils.NewId(utils.ParticipantPrefix),
		Name:      req.Name,
		RoomID:    req.RoomId,
		RoomAdmin: true,
	}

	at, err := s.authSvc.NewAccessTokenFromGrants(grants)

	if err != nil {
		return nil, twirp.NewError(twirp.Internal, "couldnt create access token")
	}

	return &twirpv1.CreateRoomResponse{AccessToken: at.ToJWT(), RoomId: req.RoomId}, nil
}

func (s *RoomService) VerifyRoom(ctx context.Context, req *twirpv1.VerifyRoomRequest) (*twirpv1.VerifyRoomResponse, error) {
	var exists bool

	if _, err := s.store.LoadRoom(ctx, req.RoomId); err == nil {
		exists = true
	}

	return &twirpv1.VerifyRoomResponse{
		Exists: exists,
	}, nil
}

func confirmExecution(f func() error) error {
	expired := time.After(APITimeoutDuration)

	for {
		select {
		case <-expired:
			return APITimeoutError
		default:
			if err := f(); err == nil {
				// only return if the execution is successfull
				return nil
			}

			time.Sleep(ConfirmExecutionInterval)
		}
	}
}
