package ping

import (
	"context"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Pingサービス
type PingService struct {
	protobuf.UnimplementedPingServer
}

// Pingサービス生成
func NewPingService() *PingService {
	return &PingService{}
}

// Pingメソッド
func (p *PingService) Ping(context.Context, *emptypb.Empty) (*protobuf.PingResponse, error) {
	return &protobuf.PingResponse{Message: "pong"}, nil
}
