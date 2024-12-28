package main

import (
	"log"
	"net"
)

func main() {
	listener, err := listen()
	if err != nil {
		log.Fatalf("Failed to start listener: %v", err)
	}

	defer func() {
		if err := listener.Close(); err != nil {
			log.Printf("Error closing listener: %v", err)
		}
	}()
}

func listen() (net.Listener, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}

	log.Printf("Listening on %s", listener.Addr())

	// Accept connections in a separate goroutine
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("Error accepting connection: %v", err)
				continue
			}

			go handleConnection(conn)
		}
	}()

	return listener, nil
}

// handleConnection handles the lifecycle of a connection
func handleConnection(conn net.Conn) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Error closing connection: %v", err)
		}
	}()
	log.Printf("Connection established with %s", conn.RemoteAddr())
}
