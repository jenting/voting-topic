package backend

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/jenting/voting-topic/backend/apis"
	"github.com/jenting/voting-topic/frontend"
)

// StartServer starts backend server
func StartServer(signalCh <-chan os.Signal) {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := apis.SetupRouter()
	frontend.SetupFrontend(router)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: router,
	}

	glog.Infof("Start server at port: %v", port)

	go func() {
		// serve connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			glog.Errorf("Shutting down the APIServer: %v", err)
		}
	}()

	// To gracefully stop all services
	<-signalCh
	glog.Infof("Shutdown server ...")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 1 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("server shutdown: %v", err)
	}
	glog.Info("Shutdown server Done")
}
