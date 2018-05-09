package cache

import (
	"sort"

	"github.com/google/uuid"
)

// Topic defines the topic voting information (upvote/downvote)
type Topic struct {
	name     string
	upvote   uint64
	downvote uint64
}

// Keeps the topics in-memory data cache
// Key: topic id ; Value: topic
var topicKV map[uuid.UUID]*Topic

func init() {
	topicKV = make(map[uuid.UUID]*Topic)
}

// CreateTopic creates a new topic
func CreateTopic(topicName string) (uuid.UUID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	topicKV[uid] = &Topic{name: topicName}
	return uid, nil
}

// GetTopic get topic accords uuid
func GetTopic(uid uuid.UUID) (*Topic, bool) {
	if v, ok := topicKV[uid]; ok {
		return v, true
	}
	return nil, false
}

// DeleteTopic deletes a topic
func DeleteTopic(uid uuid.UUID) bool {
	if _, ok := topicKV[uid]; !ok {
		// Not exists
		return true
	}

	delete(topicKV, uid)
	return true
}

// GetTopicName gets topic name
func GetTopicName(uid uuid.UUID) string {
	if v, ok := topicKV[uid]; ok {
		return v.name
	}
	return ""
}

// GetTopicUpvote gets topic upvote counts
func GetTopicUpvote(uid uuid.UUID) uint64 {
	if v, ok := topicKV[uid]; ok {
		return v.upvote
	}
	return 0
}

// GetTopicDownvote gets topic downvote counts
func GetTopicDownvote(uid uuid.UUID) uint64 {
	if v, ok := topicKV[uid]; ok {
		return v.downvote
	}
	return 0
}

// SetTopicUpvote sets topic upvote counts
func SetTopicUpvote(uid uuid.UUID, count uint64) bool {
	if v, ok := topicKV[uid]; ok {
		v.upvote = count
		return true
	}
	return false
}

// SetTopicDownvote sets topic downvote counts
func SetTopicDownvote(uid uuid.UUID, count uint64) bool {
	if v, ok := topicKV[uid]; ok {
		v.downvote = count
		return true
	}
	return false
}

// TopicListUpvote defines the topic array with upvote
type TopicListUpvote []Topic

func (l TopicListUpvote) Len() int           { return len(l) }
func (l TopicListUpvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListUpvote) Less(i, j int) bool { return l[i].upvote < l[j].upvote }

// TopicListDownvote defines the topic array with downvote
type TopicListDownvote []Topic

func (l TopicListDownvote) Len() int           { return len(l) }
func (l TopicListDownvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListDownvote) Less(i, j int) bool { return l[i].downvote < l[j].downvote }

// GetTopicDescendUpvote gets topics with desceding upvote order
func GetTopicDescendUpvote() TopicListUpvote {
	// Transfer map to array
	uvList := make(TopicListUpvote, len(topicKV))
	// Variable default value is 0
	var index int
	for _, v := range topicKV {
		uvList[index] = Topic{name: v.name, upvote: v.upvote, downvote: v.downvote}
		index++
	}

	sort.Sort(sort.Reverse(TopicListUpvote(uvList)))
	return uvList
}

// GetTopicDescendDownvote gets topics with desceding downvote order
func GetTopicDescendDownvote() TopicListDownvote {
	// Transfer map to array
	dvList := make(TopicListDownvote, len(topicKV))
	// Variable default value is 0
	var index int
	for _, v := range topicKV {
		dvList[index] = Topic{name: v.name, upvote: v.upvote, downvote: v.downvote}
		index++
	}

	sort.Sort(sort.Reverse(TopicListDownvote(dvList)))
	return dvList
}
