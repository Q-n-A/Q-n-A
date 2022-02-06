//go:build wireinject
// +build wireinject

//go:generate wire

package cmd

import (
	"github.com/Q-n-A/Q-n-A/client"
	"github.com/Q-n-A/Q-n-A/client/traq_bot"
	"github.com/Q-n-A/Q-n-A/repository"
	"github.com/Q-n-A/Q-n-A/repository/gorm2"
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/Q-n-A/Q-n-A/server/ping"
	"github.com/Q-n-A/Q-n-A/server/protobuf"
	"github.com/Q-n-A/Q-n-A/util/logger"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	provideTraQBotClientConfig,
	traq_bot.NewTraQBotClient,
	wire.Bind(new(client.BotClient), new(*traq_bot.TraQBotClient)),

	provideLoggerConfig,
	logger.NewZapLogger,

	provideRepositoryConfig,
	gorm2.NewGorm2Repository,
	wire.Bind(new(repository.Repository), new(*gorm2.Gorm2Repository)),
	gorm2.GetSqlDB,

	ping.NewPingService,
	wire.Bind(new(protobuf.PingServer), new(*ping.PingService)),

	server.NewMySQLStore,
	server.NewEcho,
	server.NewGRPCServer,

	provideServerConfig,
	server.NewServer,
)

func setupServer(config *Config) (*server.Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
