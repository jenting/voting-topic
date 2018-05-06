package apis

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/hsiaoairplane/voting-topic/backend/cache"
)

const (
	maxTopicNameLen = 255
	maxTopTopics    = 20
)

// SetupRouter returns the main gin-gonic http server
func SetupRouter() *gin.Engine {
	// Disable debug mode of gin framework.
	gin.SetMode(gin.ReleaseMode)

	// Disable console color.
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	// Create routes
	router.GET("/toptopic", getTopTopic) // get top topic
	router.GET("/topic", getTopic)       // get topic
	router.POST("/topic", createTopic)   // sumit a new topic
	router.PUT("/topic", updateTopic)    // update a topic

	return router
}

// getTopTopic returns top 20 topics (sorted by upvotes, descending)
func getTopTopic(c *gin.Context) {
	topicUpvoteDescend := cache.GetTopicDescendUpvote()
	if len(topicUpvoteDescend) > maxTopTopics {
		c.JSON(http.StatusOK, topicUpvoteDescend[:maxTopTopics])
		return
	}

	c.JSON(http.StatusOK, topicUpvoteDescend)
	return
}

// getTopic returns specific topic's update and downvote count
func getTopic(c *gin.Context) {
	// Get GET parameter
	topicName := c.Query("name")
	// Check input parameter
	if topicName == "" {
		glog.Error("Topic name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
		return
	}
	// Check if the topic exists
	if exist := cache.IsTopicExist(topicName); !exist {
		glog.Infof("Topic %s already exist", topicName)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Topic name not exist"})
		return
	}

	up := cache.GetTopicUpvote(topicName)
	down := cache.GetTopicDownvote(topicName)

	c.JSON(http.StatusOK, &topicInfo{Name: topicName, Upvote: up, Downvote: down})
	return
}

// createTopic implements the RESTful POST API.
func createTopic(c *gin.Context) {
	var t topicInfo
	if err := c.BindJSON(&t); err != nil {
		log.Println(t)
		glog.Infof("Invalid JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	// Topic should not exceed 255 characters.
	if len(t.Name) > maxTopicNameLen {
		glog.Errorf("Topic name length exceeds length %d", maxTopicNameLen)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
		return
	}

	// Check if the topic exists
	if exist := cache.IsTopicExist(t.Name); exist {
		glog.Infof("Topic %s already exist", t.Name)
		c.JSON(http.StatusOK, gin.H{"message": "Topic name already exist"})
		return
	}

	// Create new topic
	cache.CreateTopic(t.Name)
	// Set data
	cache.SetTopicUpvote(t.Name, t.Upvote)
	cache.SetTopicDownvote(t.Name, t.Downvote)

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
	return
}

// updateTopic implements the RESTful PUT API.
func updateTopic(c *gin.Context) {
	var t topicInfo
	if err := c.BindJSON(&t); err != nil {
		glog.Infof("Invalid JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	// Check if the topic exists
	if exist := cache.IsTopicExist(t.Name); !exist {
		glog.Infof("Topic %s is not exist", t.Name)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Topic name is not exist"})
		return
	}

	// Set data
	cache.SetTopicUpvote(t.Name, t.Upvote)
	cache.SetTopicDownvote(t.Name, t.Downvote)

	c.JSON(http.StatusOK, gin.H{"message": "OK"})
	return
}
