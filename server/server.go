package server

import (
	"log"
	"net"

	"github.com/Q-n-A/Q-n-A/server/grpc/ping"
	"github.com/Q-n-A/Q-n-A/server/ping_impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	s *grpc.Server
}

func NewServer(pingService *ping_impl.PingService) *Server {
	s := grpc.NewServer()

	ping.RegisterPingServer(s, pingService)

	reflection.Register(s)

	return &Server{s: s}
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
