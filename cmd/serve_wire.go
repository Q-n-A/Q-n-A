//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/Q-n-A/Q-n-A/repository"
	"github.com/Q-n-A/Q-n-A/repository/gorm2"
	"github.com/Q-n-A/Q-n-A/server"
	"github.com/google/wire"
)

var serverSet = wire.NewSet(
	newLogger,

	provideRepositoryConfig,
	gorm2.NewGormRepository,
	wire.Bind(new(repository.Repository), new(*gorm2.GormRepository)),
	gorm2.GetSqlDB,

	provideServerConfig,
	server.InjectServer,
)

func setupServer(config *Config) (*server.Server, error) {
	wire.Build(serverSet)
	return nil, nil
}
