package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/hsiaoairplane/voting-topic/backend"
)

func init() {
	// Default logging to console.
	flag.Set("logtostderr", "true")
}

func main() {
	// Parse flags.
	flag.Parse()

	// Create os channel to receives os interrupt
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// Start backend server
	backend.StartServer(signalCh)
}
