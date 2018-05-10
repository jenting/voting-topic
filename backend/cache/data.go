package cache

import (
	"sort"
	"sync/atomic"

	"github.com/google/uuid"
)

// Topic defines the Topic voting information for database
type Topic struct {
	UID      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Upvote   uint64    `json:"upvote"`
	Downvote uint64    `json:"downvote"`
}

// Keeps the topics in-memory data cache
// Key: Topic id ; Value: Topic
var topicKV map[uuid.UUID]*Topic

func init() {
	topicKV = make(map[uuid.UUID]*Topic)

	// Set init data
	uid1, _ := CreateTopic("I'm-Topic-1")
	uid2, _ := CreateTopic("I'm-Topic-2")
	uid3, _ := CreateTopic("I'm-Topic-3")

	// upvote=2 downvote=1
	IncTopicUpvote(uid1)
	IncTopicUpvote(uid1)
	IncTopicDownvote(uid1)

	// upvote=3 downvote=2
	IncTopicUpvote(uid2)
	IncTopicUpvote(uid2)
	IncTopicUpvote(uid2)
	IncTopicDownvote(uid2)
	IncTopicDownvote(uid2)

	// upvote=1 downvote=3
	IncTopicUpvote(uid3)
	IncTopicDownvote(uid3)
	IncTopicDownvote(uid3)
	IncTopicDownvote(uid3)
}

// CreateTopic creates a new Topic
func CreateTopic(topicName string) (uuid.UUID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	topicKV[uid] = &Topic{UID: uid, Name: topicName}
	return uid, nil
}

// GetTopic get Topic accords uuid
func GetTopic(uid uuid.UUID) (*Topic, bool) {
	if v, ok := topicKV[uid]; ok {
		return v, true
	}
	return nil, false
}

// DeleteTopic deletes a Topic
func DeleteTopic(uid uuid.UUID) bool {
	if _, ok := topicKV[uid]; !ok {
		// Not exists
		return true
	}

	delete(topicKV, uid)
	return true
}

// GetTopicName gets Topic name
func GetTopicName(uid uuid.UUID) string {
	if v, ok := topicKV[uid]; ok {
		return v.Name
	}
	return ""
}

// GetTopicUpvote gets Topic upvote counts
func GetTopicUpvote(uid uuid.UUID) uint64 {
	if v, ok := topicKV[uid]; ok {
		return v.Upvote
	}
	return 0
}

// GetTopicDownvote gets Topic downvote counts
func GetTopicDownvote(uid uuid.UUID) uint64 {
	if v, ok := topicKV[uid]; ok {
		return v.Downvote
	}
	return 0
}

// IncTopicUpvote sets Topic upvote counts
func IncTopicUpvote(uid uuid.UUID) bool {
	if v, ok := topicKV[uid]; ok {
		atomic.AddUint64(&v.Upvote, 1)
		return true
	}
	return false
}

// IncTopicDownvote sets Topic downvote counts
func IncTopicDownvote(uid uuid.UUID) bool {
	if v, ok := topicKV[uid]; ok {
		atomic.AddUint64(&v.Downvote, 1)
		return true
	}
	return false
}

// TopicListUpvote defines the Topic array with upvote
type TopicListUpvote []Topic

func (l TopicListUpvote) Len() int           { return len(l) }
func (l TopicListUpvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListUpvote) Less(i, j int) bool { return l[i].Upvote < l[j].Upvote }

// TopicListDownvote defines the Topic array with downvote
type TopicListDownvote []Topic

func (l TopicListDownvote) Len() int           { return len(l) }
func (l TopicListDownvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListDownvote) Less(i, j int) bool { return l[i].Downvote < l[j].Downvote }

// GetTopicDescendUpvote gets topics with desceding upvote order
func GetTopicDescendUpvote() TopicListUpvote {
	// Transfer map to array
	uvList := make(TopicListUpvote, len(topicKV))
	// Variable default value is 0
	var index int
	for k, v := range topicKV {
		uvList[index] = Topic{UID: k, Name: v.Name, Upvote: v.Upvote, Downvote: v.Downvote}
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
	for k, v := range topicKV {
		dvList[index] = Topic{UID: k, Name: v.Name, Upvote: v.Upvote, Downvote: v.Downvote}
		index++
	}

	sort.Sort(sort.Reverse(TopicListDownvote(dvList)))
	return dvList
}
