package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/mhope-2/go_networking/conn/client"
	"github.com/mhope-2/go_networking/conn/server"
)

func main() {
	done := make(chan string)
	quit := make(chan struct{})
	interrupt := make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt)

	var wg sync.WaitGroup

	// Start the listener in a goroutine
	wg.Add(1)
	go server.Listen("127.0.0.1:0", done, quit, &wg)

	// Wait for the listener to signal readiness
	listenerAddress := <-done

	// Start the dialer
	wg.Add(1)
	go client.DialAndSend(listenerAddress, &wg, 5*time.Second)

	go func() {
		<-interrupt
		log.Println("Received interrupt signal. Shutting down...")
		close(quit)
	}()

	// Wait for the dialer to finish
	wg.Wait()

	// Stop the listener
	close(quit)
	fmt.Println("Program completed")
}
