package ping

import (
	"context"

	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingService struct {
	protobuf.UnimplementedPingServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (p *PingService) Ping(context.Context, *emptypb.Empty) (*protobuf.PingResponse, error) {
	return &protobuf.PingResponse{Message: "pong"}, nil
}
