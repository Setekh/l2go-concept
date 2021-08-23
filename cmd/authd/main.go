package main

import (
	"l2go-concept/domain/auth"
	"sync"
)

var waitGroup = sync.WaitGroup{}

func main() {
	// Load database
	// Start daemon

	// Intercept clients
	startClientService()

	waitGroup.Wait()
}

func startClientService() {
	waitGroup.Add(1)
	go func() {
		defer waitGroup.Done()

		auth.StartClientServer()
	}()
}
