package main

import (
	"context"
	"flag"

	// "fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	timeout := flag.Duration("timeout", 10*time.Second, "Timeout connection")
	flag.Parse()

	host := os.Args[len(os.Args)-2]
	port := os.Args[len(os.Args)-1]

	client := NewTelnetClient(net.JoinHostPort(host,port), *timeout, os.Stdin, os.Stdout)
	err := client.Connect()
	if err != nil {
		log.Printf("Error ocures during connection to server: %s", err)
		os.Exit(1)
	} else {
		log.Println("Client connected")
	}

	notifyCtx,_ :=  signal.NotifyContext(context.Background(), os.Interrupt)

	errCh := make(chan error, 1)

	go func(client TelnetClient){ 
		if err := client.Send(); err != nil {
			log.Printf("Send error: %s", err)
			errCh <- err
		}
	}(client)

	go func(client TelnetClient){ 
		if err := client.Receive(); err != nil {
			log.Printf("Reciev error: %s", err)
			errCh <- err
		}
	}(client)


	select {
	case <-notifyCtx.Done():
		log.Printf("Closing connection...")
	case <- errCh:
		log.Printf("...EOF")
		client.Close()
	}
}