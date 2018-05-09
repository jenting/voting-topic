package cache

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateTopic(t *testing.T) {
	// Test repeat create topic
	uid, err := CreateTopic("1")
	assert.Equal(t, nil, err, "Create topic failed")

	_, exist := GetTopic(uid)
	assert.Equal(t, true, exist, "The topic should exist")
}

func TestDeleteTopic(t *testing.T) {
	uid, err := CreateTopic("2")
	assert.Equal(t, nil, err, "Create topic failed")

	ok := DeleteTopic(uid)
	assert.Equal(t, true, ok, "Delete topic failed")
}

func TestGetTopicUpvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	vote := GetTopicUpvote(uid)
	assert.EqualValues(t, 0, vote, "The upvote should be zero")

	uid, err = CreateTopic("3")
	assert.Equal(t, nil, err, "Create topic failed")

	ok := SetTopicUpvote(uid, 1)
	assert.Equal(t, true, ok, "Set topic upvote failed")

	vote = GetTopicUpvote(uid)
	assert.EqualValues(t, 1, vote, "The upvote should be one")
}

func TestGetTopicDownvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	vote := GetTopicDownvote(uid)
	assert.EqualValues(t, 0, vote, "The downvote should be zero")

	uid, err = CreateTopic("4")
	assert.Equal(t, nil, err, "Create topic failed")

	ok := SetTopicDownvote(uid, 1)
	assert.Equal(t, true, ok, "Set topic downvote failed")

	vote = GetTopicDownvote(uid)
	assert.EqualValues(t, 1, vote, "The downvote should be one")
}

func TestSetTopicUpvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	ok := SetTopicUpvote(uid, 100)
	assert.Equal(t, false, ok, "Set topic upvote should failed")

	uid, err = CreateTopic("5")
	assert.Equal(t, nil, err, "Create topic failed")

	ok = SetTopicUpvote(uid, 100)
	assert.Equal(t, true, ok, "Set topic upvote failed")

	vote := GetTopicUpvote(uid)
	assert.EqualValues(t, 100, vote, "The upvote should be one-hundred")
}

func TestSetTopicDownvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	ok := SetTopicDownvote(uid, 100)
	assert.Equal(t, false, ok, "Set topic downvote should failed")

	uid, err = CreateTopic("6")
	assert.Equal(t, nil, err, "Create topic failed")

	ok = SetTopicDownvote(uid, 100)
	assert.Equal(t, true, ok, "Set topic downvote failed")

	vote := GetTopicDownvote(uid)
	assert.EqualValues(t, 100, vote, "The downvote should be one-hundred")
}

func TestGetTopicDescendUpvote(t *testing.T) {
	tests := make(TopicListUpvote, 5)
	tests[0] = Topic{"7-5", 9, 0}
	tests[1] = Topic{"7-2", 7, 0}
	tests[2] = Topic{"7-1", 5, 0}
	tests[3] = Topic{"7-3", 3, 0}
	tests[4] = Topic{"7-4", 1, 0}

	for idx := range tests {
		uid, err := CreateTopic(tests[idx].name)
		assert.Equal(t, nil, err, "Create topic failed")

		ok := SetTopicUpvote(uid, tests[idx].upvote)
		assert.Equal(t, true, ok, "Set topic upvote failed")
	}

	topicListUpvote := GetTopicDescendUpvote()
	assert.ObjectsAreEqual(topicListUpvote, tests)
}

func TestGetTopicDescendDownvote(t *testing.T) {
	tests := make(TopicListDownvote, 5)
	tests[0] = Topic{"8-5", 0, 9}
	tests[1] = Topic{"8-2", 0, 7}
	tests[2] = Topic{"8-1", 0, 5}
	tests[3] = Topic{"8-3", 0, 3}
	tests[4] = Topic{"8-4", 0, 1}

	for idx := range tests {
		uid, err := CreateTopic(tests[idx].name)
		assert.Equal(t, nil, err, "Create topic failed")

		ok := SetTopicDownvote(uid, tests[idx].downvote)
		assert.Equal(t, true, ok, "Set topic downvote failed")
	}

	topicListDownvote := GetTopicDescendDownvote()
	assert.ObjectsAreEqual(topicListDownvote, tests)
}
