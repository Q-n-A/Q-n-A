//go:build wireinject
// +build wireinject

package server

import (
	"database/sql"

	"github.com/Q-n-A/Q-n-A/server/ping"
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var serverSet = wire.NewSet(
	ping.NewPingService,
	wire.Bind(new(protobuf.PingServer), new(*ping.PingService)),

	newMySQLStore,
	newServer,
)

func InjectServer(config *Config, db *sql.DB, logger *zap.Logger) (*Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
