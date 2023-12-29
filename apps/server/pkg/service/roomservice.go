package service

import (
	"context"
	"errors"
	"sync/atomic"
	"time"

	"github.com/acerohernan/meet/core"
	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/config/logger"
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

func confirmExecution(f func() error) error {
	counter := atomic.Int32{}
	expired := time.After(APITimeoutDuration)

	for {
		select {
		case <-expired:
			return APITimeoutError
		default:
			err := f()

			counter.Add(1)
			// only return if the execution is successfull
			if err == nil {
				logger.Infow("function finished", "executed", counter.Load())
				return nil
			}

			time.Sleep(ConfirmExecutionInterval)
		}
	}
}
