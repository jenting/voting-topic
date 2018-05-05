package cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsTopicExist(t *testing.T) {
	exist := IsTopicExist("1")
	assert.Equal(t, false, exist, "The topic should not exist")
}

func TestCreateTopic(t *testing.T) {
	ok := CreateTopic("2")
	assert.Equal(t, true, ok, "Create topic failed")

	// Test repeat create topic
	ok = CreateTopic("2")
	assert.Equal(t, true, ok, "Create topic failed")

	exist := IsTopicExist("2")
	assert.Equal(t, true, exist, "The topic should exist")
}

func TestDeleteTopic(t *testing.T) {
	ok := DeleteTopic("3")
	assert.Equal(t, true, ok, "Delete topic success")

	ok = DeleteTopic("3")
	assert.Equal(t, true, ok, "Delete topic failed")

	ok = CreateTopic("3")
	assert.Equal(t, true, ok, "Create topic failed")

	ok = DeleteTopic("3")
	assert.Equal(t, true, ok, "Delete topic failed")

	exist := IsTopicExist("3")
	assert.Equal(t, false, exist, "The topic should not exist")
}

func TestGetTopicUpvote(t *testing.T) {
	vote := GetTopicUpvote("4")
	assert.EqualValues(t, 0, vote, "The upvote should be zero")

	ok := CreateTopic("4")
	assert.Equal(t, true, ok, "Create topic failed")

	ok = SetTopicUpvote("4", 1)
	assert.Equal(t, true, ok, "Set topic upvote failed")

	vote = GetTopicUpvote("4")
	assert.EqualValues(t, 1, vote, "The upvote should be one")
}

func TestGetTopicDownvote(t *testing.T) {
	vote := GetTopicDownvote("5")
	assert.EqualValues(t, 0, vote, "The downvote should be zero")

	ok := CreateTopic("5")
	assert.Equal(t, true, ok, "Create topic failed")

	ok = SetTopicDownvote("5", 1)
	assert.Equal(t, true, ok, "Set topic downvote failed")

	vote = GetTopicDownvote("5")
	assert.EqualValues(t, 1, vote, "The downvote should be one")
}

func TestSetTopicUpvote(t *testing.T) {
	ok := SetTopicUpvote("6", 100)
	assert.Equal(t, false, ok, "Set topic upvote should failed")

	ok = CreateTopic("6")
	assert.Equal(t, true, ok, "Create topic failed")

	ok = SetTopicUpvote("6", 100)
	assert.Equal(t, true, ok, "Set topic upvote failed")

	vote := GetTopicUpvote("6")
	assert.EqualValues(t, 100, vote, "The upvote should be one-hundred")
}

func TestSetTopicDownvote(t *testing.T) {
	ok := SetTopicDownvote("7", 100)
	assert.Equal(t, false, ok, "Set topic downvote should failed")

	ok = CreateTopic("7")
	assert.Equal(t, true, ok, "Create topic failed")

	ok = SetTopicDownvote("7", 100)
	assert.Equal(t, true, ok, "Set topic downvote failed")

	vote := GetTopicDownvote("7")
	assert.EqualValues(t, 100, vote, "The downvote should be one-hundred")
}

func TestGetTopicDescendUpvote(t *testing.T) {
	tests := make(TopicListUpvote, 5)
	tests[0] = topic{"8-5", 9, 0}
	tests[1] = topic{"8-2", 7, 0}
	tests[2] = topic{"8-1", 5, 0}
	tests[3] = topic{"8-3", 3, 0}
	tests[4] = topic{"8-4", 1, 0}

	for idx := range tests {
		ok := CreateTopic(tests[idx].name)
		assert.Equal(t, true, ok, "Create topic failed")

		ok = SetTopicUpvote(tests[idx].name, tests[idx].upvote)
		assert.Equal(t, true, ok, "Set topic upvote failed")
	}

	topicListUpvote := GetTopicDescendUpvote()
	assert.ObjectsAreEqual(topicListUpvote, tests)
}

func TestGetTopicDescendDownvote(t *testing.T) {
	tests := make(TopicListDownvote, 5)
	tests[0] = topic{"9-5", 0, 9}
	tests[1] = topic{"9-2", 0, 7}
	tests[2] = topic{"9-1", 0, 5}
	tests[3] = topic{"9-3", 0, 3}
	tests[4] = topic{"9-4", 0, 1}

	for idx := range tests {
		ok := CreateTopic(tests[idx].name)
		assert.Equal(t, true, ok, "Create topic failed")

		ok = SetTopicDownvote(tests[idx].name, tests[idx].downvote)
		assert.Equal(t, true, ok, "Set topic downvote failed")
	}

	topicListDownvote := GetTopicDescendDownvote()
	assert.ObjectsAreEqual(topicListDownvote, tests)
}
