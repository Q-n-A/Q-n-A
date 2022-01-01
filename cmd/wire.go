//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/Q-n-A/Q-n-A/server/ping"
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var serverSet = wire.NewSet(
	ping.NewPingService,
	wire.Bind(new(protobuf.PingServer), new(*ping.PingService)),

	provideServerConfig,
	server.NewServer,
)

func SetupServer(config *Config, logger *zap.Logger) (*server.Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
