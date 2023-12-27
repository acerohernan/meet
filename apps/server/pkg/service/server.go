package service

import (
	"fmt"
	"net"
	"net/http"

	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/rs/cors"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	conf       *config.Config
	httpServer *http.Server
	doneChan   chan struct{}
}

func NewServer(conf *config.Config, roomSvc *RoomService) *Server {
	mux := http.NewServeMux()

	roomServer := twirpv1.NewRoomServiceServer(roomSvc)
	mux.Handle(roomServer.PathPrefix(), roomServer)

	// health check
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// CORS is allowed, the authentication is made with JWT
	handler := cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		AllowedHeaders: []string{"*"},
		// allow preflight to be cached for a day
		MaxAge: 86400,
	})

	return &Server{
		conf:     conf,
		doneChan: make(chan struct{}),
		httpServer: &http.Server{
			Handler: handler.Handler(mux),
		},
	}
}

func (s *Server) Start() error {
	httpGroup := &errgroup.Group{}

	l, err := net.Listen("tcp", fmt.Sprint(":", s.conf.Port))

	if err != nil {
		return err
	}

	httpGroup.Go(func() error {
		return s.httpServer.Serve(l)
	})

	// listen for errors in http server
	go func() {
		if err := httpGroup.Wait(); err != http.ErrServerClosed {
			logger.Errow("could not start the server: ", err)
			s.Stop()
		}
	}()

	logger.Infow("http server running!", "url", fmt.Sprint("http://localhost:", s.conf.Port))

	<-s.doneChan

	return nil
}

func (s *Server) Stop() error {
	if err := s.httpServer.Close(); err != nil {
		logger.Errow("Error at closing http server", err)
	}

	close(s.doneChan)

	return nil
}
