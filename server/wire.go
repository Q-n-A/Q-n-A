//go:build wireinject
// +build wireinject

package server

import (
	"github.com/Q-n-A/Q-n-A/server/ping_impl"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	ping_impl.NewPingService,
	NewServer,
)

func InjectServer() (*Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
