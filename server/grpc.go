package server

import (
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// 新しいgRPCサーバーを生成
func newGRPCServer(logger *zap.Logger, pingService protobuf.PingServer) *grpc.Server {
	// loggerを使いgRPCサーバーを生成
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(logger),
		)),
	)

	// reflectionを有効にする
	reflection.Register(s)

	// サービスを登録
	registerServices(s, pingService)

	return s
}

// サービスを登録
func registerServices(s *grpc.Server, pingService protobuf.PingServer) {
	protobuf.RegisterPingServer(s, pingService)
}
