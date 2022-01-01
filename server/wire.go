//go:build wireinject
// +build wireinject

package server

import (
	"github.com/Q-n-A/Q-n-A/server/ping"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var serverSet = wire.NewSet(
	ping.NewPingService,
	NewServer,
)

func InjectServer(*zap.Logger) (*Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
