package backend

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/hsiaoairplane/voting-topic/backend/apis"
)

var (
	port uint
)

func init() {
	flag.UintVar(&port, "port", 8080, "Server port should assigned")
}

// StartServer starts backend server
func StartServer(router *gin.Engine, signalCh <-chan os.Signal) {
	router.GET("/toptopic", apis.GetTopTopic) // get top topic
	router.GET("/topic", apis.GetTopic)       // get topic
	router.POST("/topic", apis.CreateTopic)   // sumit a new topic
	router.PUT("/topic", apis.UpdateTopic)    // update a topic

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
	// a timeout of 10 seconds.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		glog.Fatalf("server shutdown: %v", err)
	}
	glog.Info("Shutdown server Done")
}
