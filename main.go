package main

import (
	"log"

	"github.com/Q-n-A/Q-n-A/server"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()

	logger.Info("Q'n'A - traP Anonymous Question Box Service")

	s, err := server.InjectServer(logger)
	if err != nil {
		log.Panicf("failed to setup gRPC server: %v", err)
	}

	s.Run()
}
