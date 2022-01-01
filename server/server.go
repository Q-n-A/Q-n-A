package server

import (
	"fmt"
	"net"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Server struct {
	e      *echo.Echo
	s      *grpc.Server
	logger *zap.Logger
	c      *Config
}

type Config struct {
	GRPCPort int
	RESTPort int
}

func newServer(Config *Config, logger *zap.Logger, store sessions.Store, pingService protobuf.PingServer) *Server {
	s := newGRPCServer(logger)
	setupServices(s, pingService)

	e := newEcho(store)
	setupHandlers(e, s)

	return &Server{
		e:      e,
		s:      s,
		logger: logger,
		c:      Config,
	}
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.c.GRPCPort))
	if err != nil {
		s.logger.Panic("failed to setup Listener", zap.Error(err))
	}

	s.logger.Info("starting gRPC server on port " + fmt.Sprintf("%d", s.c.GRPCPort))

	go func() {
		if err := s.s.Serve(lis); err != nil {
			s.logger.Panic("failed to run gRPC server", zap.Error(err))
		}
	}()

	s.e.Logger.Panic(s.e.Start(fmt.Sprintf(":%d", s.c.RESTPort)))
}
