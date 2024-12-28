package main

import (
	"fmt"
	"github.com/mhope-2/go_networking/conn/client"
	"github.com/mhope-2/go_networking/conn/server"
	"sync"
	"time"
)

func main() {
	done := make(chan string)
	quit := make(chan struct{})
	var wg sync.WaitGroup

	// Start the listener in a goroutine
	wg.Add(1)
	go server.Listen("127.0.0.1:0", done, quit, &wg)

	// Wait for the listener to signal readiness
	listenerAddress := <-done

	// Start the dialer
	wg.Add(1)
	go client.DialAndSend(listenerAddress, &wg, 5*time.Second)

	// Wait for the dialer to finish
	wg.Wait()

	// Stop the listener
	close(quit)
	fmt.Println("Program completed")
}
