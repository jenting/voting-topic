package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"

	"github.com/hsiaoairplane/voting-topic/backend"
)

func init() {
	// Default logging to console.
	flag.Set("logtostderr", "true")
}

func main() {
	// Parse flags.
	flag.Parse()

	// Disable debug mode of gin framework.
	gin.SetMode(gin.ReleaseMode)

	// Disable console color.
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Create os channel to receives os interrupt
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// Start backend server
	backend.StartServer(router, signalCh)
}
