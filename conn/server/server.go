package server

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// Listen starts a TCP server and signals when it's ready.
func Listen(address string, done chan string, quit chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// Create a listener
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error starting listener: %v", err)
	}
	defer listener.Close()
	fmt.Printf("Listener started at: %s\n", listener.Addr())

	// Signal readiness to the main function
	done <- listener.Addr().String()

	// Accept connections until quit is signaled
	for {
		select {
		case <-quit:
			fmt.Println("Shutting down listener")
			return
		default:
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				return
			}

			go handleConnection(conn)
		}
	}
}

// HandleConnection handles incoming connections.
func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Printf("Error reading from connection: %v", err)
			}
			return
		}
		fmt.Printf("Received: %q\n", buf[:n])
	}
}
