package main

import (
	"l2go-concept/domain/game"
	"l2go-concept/domain/game/storage"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	// Load database
	store := storage.CreateStorage()

	// Start daemon

	// Intercept clients
	startClientService(store)

	waitGroup.Wait()
}

func startClientService(store storage.GameStorage) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		game.StartClientServer(store)
	}()
}
