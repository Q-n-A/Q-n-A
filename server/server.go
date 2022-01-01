package server

import (
	"fmt"
	"log"
	"net"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	s      *grpc.Server
	logger *zap.Logger
	c      *Config
}

type Config struct {
	GRPCPort int
	RESTPort int
}

func NewServer(logger *zap.Logger, Config *Config, pingServer protobuf.PingServer) *Server {
	s := newGRPCServer(logger)

	protobuf.RegisterPingServer(s, pingServer)

	return &Server{
		s:      s,
		logger: logger,
		c:      Config,
	}
}

func newGRPCServer(logger *zap.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	reflection.Register(s)

	return s
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.c.GRPCPort))
	if err != nil {
		log.Panicf("failed to setup Listener: %v", err)
	}

	s.logger.Info("starting gRPC server on port " + fmt.Sprintf("%d", s.c.GRPCPort))

	if err := s.s.Serve(lis); err != nil {
		log.Panicf("failed to run gRPC server: %v", err)
	}
}
