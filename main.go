package main

import (
	"log"

	"github.com/Q-n-A/Q-n-A/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Panicf("failed to start Q'n'A application: %v", err)
	}
}
