package backend

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/hsiaoairplane/voting-topic/backend/apis"
	"github.com/hsiaoairplane/voting-topic/frontend"
)

var (
	port uint
)

func init() {
	flag.UintVar(&port, "port", 8080, "Server port should assigned")
}

// StartServer starts backend server
func StartServer(signalCh <-chan os.Signal) {
	router := apis.SetupRouter()
	frontend.SetupFrontend(router)

	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	glog.Infof("Start server at port: %d", port)

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
	// a timeout of 3 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("server shutdown: %v", err)
	}
	glog.Info("Shutdown server Done")
}
