package main

import (
	"l2go-concept/internal/auth/network"
	"l2go-concept/internal/auth/storage"
	auth2 "l2go-concept/pkg/auth"
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

func startClientService(store auth2.Storage) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		network.StartClientServer(store)
	}()
}
