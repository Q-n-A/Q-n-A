//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/Q-n-A/Q-n-A/repository"
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

var serverSet = wire.NewSet(
	wire.Value([]zap.Option{}),
	zap.NewProduction,

	provideRepositoryConfig,
	repository.NewGormRepository,
	wire.Bind(new(repository.Repository), new(*repository.GormRepository)),
	repository.GetSqlDB,

	provideServerConfig,
	server.InjectServer,
)

func SetupServer(config *Config) (*server.Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
