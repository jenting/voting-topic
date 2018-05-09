package cache

import (
	"sort"

	"github.com/google/uuid"
)

// Topic defines the Topic voting information for database
type Topic struct {
	UID      uuid.UUID
	Name     string
	Upvote   uint64
	Downvote uint64
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

	SetTopicUpvote(uid1, 2)
	SetTopicDownvote(uid1, 1)
	SetTopicUpvote(uid2, 3)
	SetTopicDownvote(uid2, 2)
	SetTopicUpvote(uid3, 1)
	SetTopicDownvote(uid3, 3)
}

// CreateTopic creates a new Topic
func CreateTopic(topicName string) (uuid.UUID, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return uuid.Nil, err
	}

	topicKV[uid] = &Topic{Name: topicName}
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

// SetTopicUpvote sets Topic upvote counts
func SetTopicUpvote(uid uuid.UUID, count uint64) bool {
	if v, ok := topicKV[uid]; ok {
		v.Upvote = count
		return true
	}
	return false
}

// SetTopicDownvote sets Topic downvote counts
func SetTopicDownvote(uid uuid.UUID, count uint64) bool {
	if v, ok := topicKV[uid]; ok {
		v.Downvote = count
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
