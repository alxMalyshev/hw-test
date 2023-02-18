package main

import (
	"context"
	"flag"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

var timeout = flag.Duration("timeout", 10*time.Second, "Timeout connection")

func main() {
	flag.Parse()

	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	client := NewTelnetClient(net.JoinHostPort(host, port), *timeout, os.Stdin, os.Stdout)
	if err := client.Connect(); err != nil {
		log.Printf("Error ocures during connection to server: %s", err)
		os.Exit(1)
	} else {
		log.Println("Client connected")
	}

	notifyCtx, _ := signal.NotifyContext(context.Background(), os.Interrupt)

	errCh := make(chan error, 1)

	go func(client TelnetClient) {
		errCh <- client.Send()
	}(client)

	go func(client TelnetClient) {
		errCh <- client.Receive()
	}(client)

	select {
	case <-notifyCtx.Done():
		log.Printf("Closing connection...")
	case <-errCh:
		log.Printf("...EOF")
		client.Close()
	}
}
