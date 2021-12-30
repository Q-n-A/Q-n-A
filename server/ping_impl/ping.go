package ping_impl

import (
	"context"

	"github.com/Q-n-A/Q-n-A/server/grpc/ping"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingService struct {
	ping.UnimplementedPingServer
}

func NewPingService() *PingService {
	return &PingService{}
}

func (p *PingService) Ping(context.Context, *emptypb.Empty) (*ping.PingResponse, error) {
	return &ping.PingResponse{Message: "pong"}, nil
}
