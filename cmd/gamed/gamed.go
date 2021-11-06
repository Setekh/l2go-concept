package main

import (
	"l2go-concept/internal/game"
	"l2go-concept/internal/game/network"
	"l2go-concept/internal/game/storage"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	dependencyManager := game.CreateDependencyManager(game.CreateTimeController(), storage.CreateStorage())

	// Intercept clients
	startClientService(dependencyManager)

	waitGroup.Wait()
}

func startClientService(dependencyManager game.DependencyManager) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		network.StartClientServer(dependencyManager)
	}()
}
