package service

import (
	"fmt"
	"net"
	"net/http"

	twirpv1 "github.com/acerohernan/meet/core/twirp/v1"
	"github.com/acerohernan/meet/pkg/config"
	"github.com/acerohernan/meet/pkg/config/logger"
	"github.com/acerohernan/meet/pkg/service/auth"
	"github.com/acerohernan/meet/pkg/service/router"
	"github.com/acerohernan/meet/pkg/service/rtc"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"golang.org/x/sync/errgroup"
)

type Server struct {
	conf        *config.Config
	httpServer  *http.Server
	router      *router.Router
	doneChan    chan struct{}
	rtcMananger rtc.RTCManager
}

func NewServer(conf *config.Config, authMiddleware *auth.AuthMiddleware, roomSvc *RoomService, router *router.Router, rtcManager rtc.RTCManager) *Server {
	mux := http.NewServeMux()

	// health check
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// rtc endpoints
	mux.HandleFunc("/rtc", rtcManager.ServeWS)
	mux.HandleFunc("/join", rtcManager.ServeJoinRequest)

	// twirp services
	roomServer := twirpv1.NewRoomServiceServer(roomSvc)
	mux.Handle(roomServer.PathPrefix(), roomServer)

	middlewares := []negroni.Handler{
		negroni.NewRecovery(),

		// CORS is allowed, the authentication is made with JWT
		cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowedHeaders: []string{"*"},
			// allow preflight to be cached for a day
			MaxAge: 86400,
		}),

		authMiddleware,
	}

	// apply middlewares
	handler := negroni.New()
	for _, m := range middlewares {
		handler.Use(m)
	}
	handler.UseHandler(mux)

	return &Server{
		conf:     conf,
		doneChan: make(chan struct{}),
		httpServer: &http.Server{
			Handler: handler,
		},
		router:      router,
		rtcMananger: rtcManager,
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
			logger.Errorw("could not start the server: ", err)
			s.Stop()
		}
	}()

	logger.Infow("http server running!", "url", fmt.Sprint("http://localhost:", s.conf.Port))

	// start router
	localNode, err := s.router.Start()

	if err != nil {
		return err
	}

	logger.Infow("node started successfully", "nodeID", localNode.Id)

	<-s.doneChan

	return nil
}

func (s *Server) Stop() error {
	if err := s.httpServer.Close(); err != nil {
		logger.Errorw("error at closing http server", err)
	}

	if err := s.rtcMananger.Close(); err != nil {
		logger.Errorw("error at closing rtc manager", err)
	}

	if err := s.router.Stop(); err != nil {
		logger.Errorw("error at stopping router", err)
	}

	close(s.doneChan)

	return nil
}
