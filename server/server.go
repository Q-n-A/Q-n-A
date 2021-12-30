package server

import (
	"log"
	"net"

	"github.com/Q-n-A/Q-n-A/server/grpc/ping"
	"github.com/Q-n-A/Q-n-A/server/ping_impl"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	s *grpc.Server
}

func NewServer(pingService *ping_impl.PingService) *Server {
	s := newGRPCServer()

	ping.RegisterPingServer(s, pingService)

	return &Server{s: s}
}

func newGRPCServer() *grpc.Server {
	logger, _ := zap.NewProduction()

	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	reflection.Register(s)

	return s
}

func (s *Server) Run() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panicf("failed to setup Listener: %v", err)
	}

	log.Println("starting gRPC server on port 9000")

	if err := s.s.Serve(lis); err != nil {
		log.Panicf("failed to run gRPC server: %v", err)
	}
}
