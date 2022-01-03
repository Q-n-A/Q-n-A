package server

import (
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func newGRPCServer(logger *zap.Logger) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	reflection.Register(s)

	return s
}

func setupServices(s *grpc.Server, pingService protobuf.PingServer) {
	protobuf.RegisterPingServer(s, pingService)
}
