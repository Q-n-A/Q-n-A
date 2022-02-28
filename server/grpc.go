package server

import (
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// NewGRPCServer 新しいgRPCサーバーを生成
func NewGRPCServer(logger *zap.Logger, pingService protobuf.PingServer) *grpc.Server {
	// loggerを使いgRPCサーバーを生成
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	// reflectionを有効にする
	reflection.Register(s)

	// サービスサーバーを登録
	registerServers(s, pingService)

	return s
}

// registerServers サービスサーバーを登録
func registerServers(s *grpc.Server, pingService protobuf.PingServer) {
	protobuf.RegisterPingServer(s, pingService)
}
