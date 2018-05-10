package cache

import (
	"sync"
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

	ok := IncTopicUpvote(uid)
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

	ok := IncTopicDownvote(uid)
	assert.Equal(t, true, ok, "Set topic downvote failed")

	vote = GetTopicDownvote(uid)
	assert.EqualValues(t, 1, vote, "The downvote should be one")
}

func TestSetTopicUpvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	ok := IncTopicUpvote(uid)
	assert.Equal(t, false, ok, "Set topic upvote should failed")

	uid, err = CreateTopic("5")
	assert.Equal(t, nil, err, "Create topic failed")

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok = IncTopicUpvote(uid)
			assert.Equal(t, true, ok, "Set topic upvote failed")
		}()
	}
	wg.Wait()

	vote := GetTopicUpvote(uid)
	assert.EqualValues(t, 100, vote, "The upvote should be one-hundred")
}

func TestSetTopicDownvote(t *testing.T) {
	uid, err := uuid.NewRandom()
	assert.Equal(t, nil, err, "New random failed")

	ok := IncTopicDownvote(uid)
	assert.Equal(t, false, ok, "Set topic downvote should failed")

	uid, err = CreateTopic("6")
	assert.Equal(t, nil, err, "Create topic failed")

	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ok := IncTopicDownvote(uid)
			assert.Equal(t, true, ok, "Set topic downvote failed")
		}()
	}
	wg.Wait()

	vote := GetTopicDownvote(uid)
	assert.EqualValues(t, 100, vote, "The downvote should be one-hundred")
}

func TestGetTopicDescendUpvote(t *testing.T) {
	tests := make(TopicListUpvote, 5)
	tests[0] = Topic{Name: "7-5", Upvote: 9, Downvote: 0}
	tests[1] = Topic{Name: "7-2", Upvote: 7, Downvote: 0}
	tests[2] = Topic{Name: "7-1", Upvote: 5, Downvote: 0}
	tests[3] = Topic{Name: "7-3", Upvote: 3, Downvote: 0}
	tests[4] = Topic{Name: "7-4", Upvote: 1, Downvote: 0}

	for idx := range tests {
		uid, err := CreateTopic(tests[idx].Name)
		assert.Equal(t, nil, err, "Create topic failed")

		wg := sync.WaitGroup{}
		for j := 0; j < int(tests[idx].Downvote); j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ok := IncTopicUpvote(uid)
				assert.Equal(t, true, ok, "Set topic upvote failed")
			}()
		}
		wg.Wait()
	}

	topicListUpvote := GetTopicDescendUpvote()
	assert.ObjectsAreEqual(topicListUpvote, tests)
}

func TestGetTopicDescendDownvote(t *testing.T) {
	tests := make(TopicListDownvote, 5)
	tests[0] = Topic{Name: "8-5", Upvote: 0, Downvote: 9}
	tests[1] = Topic{Name: "8-2", Upvote: 0, Downvote: 7}
	tests[2] = Topic{Name: "8-1", Upvote: 0, Downvote: 5}
	tests[3] = Topic{Name: "8-3", Upvote: 0, Downvote: 3}
	tests[4] = Topic{Name: "8-4", Upvote: 0, Downvote: 1}

	for idx := range tests {
		uid, err := CreateTopic(tests[idx].Name)
		assert.Equal(t, nil, err, "Create topic failed")

		wg := sync.WaitGroup{}
		for j := 0; j < int(tests[idx].Downvote); j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				ok := IncTopicDownvote(uid)
				assert.Equal(t, true, ok, "Set topic downvote failed")
			}()
		}
		wg.Wait()
	}

	topicListDownvote := GetTopicDescendDownvote()
	assert.ObjectsAreEqual(topicListDownvote, tests)
}
