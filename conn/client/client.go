package client

import (
	"log"
	"net"
	"sync"
	"time"
)

// DialAndSend connects to the server and sends a message.
func DialAndSend(address string, wg *sync.WaitGroup, timeout time.Duration) {
	defer wg.Done()

	conn, err := DialWithTimeout("tcp", address, timeout)
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	message := "Hello from the dialer!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Printf("Error sending message: %v", err)
	}
}

func DialWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Timeout: timeout,
	}
	return d.Dial(network, address)
}
