package frontend

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsiaoairplane/voting-topic/backend/cache"
)

// SetupFrontend setup frontend routes.
func SetupFrontend(router *gin.Engine) {
	// Create route
	router.GET("/", renderHTML)

	router.LoadHTMLFiles("./frontend/index.html")
}

func renderHTML(c *gin.Context) {
	// Display homepage
	c.HTML(http.StatusOK, "index.html",
		gin.H{
			"title":     "Hola cómo estás",
			"toptopics": cache.GetTopicDescendUpvote(),
		},
	)
}
