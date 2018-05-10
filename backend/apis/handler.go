package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/google/uuid"

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
	router.GET("/toptopic", getTopTopic)               // get top topic
	router.GET("/topic", getTopic)                     // get topic
	router.POST("/topic", createTopic)                 // sumit a new topic
	router.PUT("/topic/upvote", updateTopicUpvote)     // update topic's upvote
	router.PUT("/topic/downvote", updateTopicDownvote) // update topic's downvote

	return router
}

// getTopic returns specific topic's update and downvote count
func getTopic(c *gin.Context) {
	inputUUID := c.Query("uid")

	// Get GET parameter
	uid, err := uuid.Parse(inputUUID)
	if err != nil {
		glog.Errorf("Invalid input uid: %v", inputUUID)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input uid"})
		return
	}

	// Get topic
	_, ok := cache.GetTopic(uid)
	if ok == false {
		glog.Errorf("Get topic %v failed", uid)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Topic not exist"})
		return
	}

	name := cache.GetTopicName(uid)
	up := cache.GetTopicUpvote(uid)
	down := cache.GetTopicDownvote(uid)

	c.JSON(http.StatusOK, &cache.Topic{UID: uid, Name: name, Upvote: up, Downvote: down})
	return
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

// createTopic implements the RESTful POST API.
func createTopic(c *gin.Context) {
	var t cache.Topic
	if err := c.ShouldBindJSON(&t); err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	// Topic should not exceed 255 characters.
	if len(t.Name) > maxTopicNameLen {
		glog.Errorf("Topic name length exceeds length %d", maxTopicNameLen)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Topic name over length"})
		return
	}

	// Create new topic
	uid, err := cache.CreateTopic(t.Name)
	if err != nil {
		glog.Errorf("Create topic %v err: %v", t.Name, err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Create topic failed"})
		return
	}

	topic, _ := cache.GetTopic(uid)

	c.JSON(http.StatusOK, topic)
	return
}

// updateTopicUpvote implements the RESTful PUT API.
func updateTopicUpvote(c *gin.Context) {
	var t cache.Topic
	if err := c.ShouldBindJSON(&t); err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	_, ok := cache.GetTopic(t.UID)
	if ok == false {
		glog.Errorf("UUID %v not exist", t.UID)
		c.JSON(http.StatusBadRequest, gin.H{"message": "UUID not exist"})
		return
	}

	// Set data
	_ = cache.IncTopicUpvote(t.UID)
	topic, _ := cache.GetTopic(t.UID)

	c.JSON(http.StatusOK, topic)
	return
}

// updateTopicDownvote implements the RESTful PUT API.
func updateTopicDownvote(c *gin.Context) {
	var t cache.Topic
	if err := c.ShouldBindJSON(&t); err != nil {
		glog.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	_, ok := cache.GetTopic(t.UID)
	if ok == false {
		glog.Errorf("UUID %v not exist", t.UID)
		c.JSON(http.StatusBadRequest, gin.H{"message": "UUID not exist"})
		return
	}

	// Set data
	_ = cache.IncTopicDownvote(t.UID)
	topic, _ := cache.GetTopic(t.UID)

	c.JSON(http.StatusOK, topic)
	return
}
