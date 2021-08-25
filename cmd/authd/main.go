package main

import (
	"l2go-concept/domain/auth"
	"l2go-concept/domain/auth/storage"
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

func startClientService(store storage.LoginStorage) {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		auth.StartClientServer(store)
	}()
}
