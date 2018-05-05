package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/hsiaoairplane/voting-topic/backend/cache"
)

const (
	maxTopicNameLen = 255
	maxTopTopics    = 20
)

// GetTopTopic returns top 20 topics (sorted by upvotes, descending)
func GetTopTopic(c *gin.Context) {
	topicUpvoteDescend := cache.GetTopicDescendUpvote()
	if len(topicUpvoteDescend) > maxTopTopics {
		c.JSON(http.StatusOK, topicUpvoteDescend[:maxTopTopics])
		return
	}

	c.JSON(http.StatusOK, topicUpvoteDescend)
	return
}

// GetTopic returns specific topic's update and downvote count
func GetTopic(c *gin.Context) {
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

// CreateTopic implements the RESTful POST API.
func CreateTopic(c *gin.Context) {
	var t topicInfo
	if err := c.BindJSON(&t); err != nil {
		glog.Infof("Invalid JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	// Check input parameter
	if t.Name == "" {
		glog.Error("Topic name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
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

// UpdateTopic implements the RESTful PUT API.
func UpdateTopic(c *gin.Context) {
	var t topicInfo
	if err := c.BindJSON(&t); err != nil {
		glog.Infof("Invalid JSON input")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON parameter"})
		return
	}

	// Check input parameter
	if t.Name == "" {
		glog.Error("Topic name is empty")
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid parameter"})
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
