package main

import (
	"fmt"
	"log"

	"github.com/Q-n-A/Q-n-A/server"
)

func main() {
	fmt.Println("Q'n'A - traP Anonymous Question Box Service")

	s, err := server.InjectServer()
	if err != nil {
		log.Panicf("failed to setup gRPC server: %v", err)
	}

	s.Run()
}
