package ping

import (
	"context"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server Pingサーバー
type Server struct {
	protobuf.UnimplementedPingServer
}

// NewServer Pingサーバー生成
func NewServer() *Server {
	return &Server{}
}

// Ping Pingメソッド
func (p *Server) Ping(context.Context, *emptypb.Empty) (*protobuf.PingResponse, error) {
	return &protobuf.PingResponse{Message: "pong"}, nil
}
