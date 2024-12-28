package main

import (
	"net"
	"testing"
)

func testListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	// close listener
	defer func() {
		if err := listener.Close(); err != nil {
			t.Fatalf("Failed to close listener with err: %s", err)
		}
	}()

	t.Logf("bound to %q", listener.Addr())
}
