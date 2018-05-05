package cache

import "sort"

// topic defines the topic voting information (upvote/downvote)
type vote struct {
	upvote   uint64
	downvote uint64
}

// Keeps the topics in-memory data cache
// Key: topic name ; Value: vote
var topicKV map[string]vote

func init() {
	topicKV = make(map[string]vote)
}

// Sorting via criteria upvote
type topic struct {
	name     string
	upvote   uint64
	downvote uint64
}

// TopicListUpvote defines the topic array with upvote
type TopicListUpvote []topic

func (l TopicListUpvote) Len() int           { return len(l) }
func (l TopicListUpvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListUpvote) Less(i, j int) bool { return l[i].upvote < l[j].upvote }

// TopicListDownvote defines the topic array with downvote
type TopicListDownvote []topic

func (l TopicListDownvote) Len() int           { return len(l) }
func (l TopicListDownvote) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l TopicListDownvote) Less(i, j int) bool { return l[i].downvote < l[j].downvote }

// GetTopicDescendUpvote gets topics with desceding upvote order
func GetTopicDescendUpvote() TopicListUpvote {
	// Transfer map to array
	uvList := make(TopicListUpvote, len(topicKV))
	// Variable default value is 0
	var index int
	for k, v := range topicKV {
		uvList[index] = topic{name: k, upvote: v.upvote, downvote: v.downvote}
		index++
	}

	sort.Sort(TopicListUpvote(uvList))
	return uvList
}

// GetTopicDescendDownvote gets topics with desceding downvote order
func GetTopicDescendDownvote() TopicListDownvote {
	// Transfer map to array
	dvList := make(TopicListDownvote, len(topicKV))
	// Variable default value is 0
	var index int
	for k, v := range topicKV {
		dvList[index] = topic{name: k, upvote: v.upvote, downvote: v.downvote}
		index++
	}

	sort.Sort(TopicListDownvote(dvList))
	return dvList
}

// IsTopicExist checks is the topic name exists in in-memory data cache
func IsTopicExist(name string) bool {
	_, ok := topicKV[name]
	return ok
}

// CreateTopic creates a new topic
func CreateTopic(name string) bool {
	if _, ok := topicKV[name]; ok {
		// Already exists
		return true
	}

	topicKV[name] = vote{}
	return true
}

// DeleteTopic deletes a topic
func DeleteTopic(name string) bool {
	if _, ok := topicKV[name]; !ok {
		// Already not exists
		return true
	}

	delete(topicKV, name)
	return true
}

// GetTopicUpvote gets topic upvote counts
func GetTopicUpvote(name string) uint64 {
	if v, ok := topicKV[name]; ok {
		return v.upvote
	}
	return 0
}

// GetTopicDownvote gets topic downvote counts
func GetTopicDownvote(name string) uint64 {
	if v, ok := topicKV[name]; ok {
		return v.downvote
	}
	return 0
}

// SetTopicUpvote sets topic upvote counts
func SetTopicUpvote(name string, count uint64) bool {
	if v, ok := topicKV[name]; ok {
		v.upvote = count
		return true
	}
	return false
}

// SetTopicDownvote sets topic downvote counts
func SetTopicDownvote(name string, count uint64) bool {
	if v, ok := topicKV[name]; ok {
		v.downvote = count
		return true
	}
	return false
}

// IncTopicUpvote increases topic upvote counts
func IncTopicUpvote(name string) bool {
	if v, ok := topicKV[name]; ok {
		v.upvote = v.upvote + 1
		return true
	}
	return false
}

// IncTopicDownvote increases topic downvote counts
func IncTopicDownvote(name string) bool {
	if v, ok := topicKV[name]; ok {
		v.downvote = v.downvote + 1
		return true
	}
	return false
}

// DecTopicUpvote decreases topic upvote counts
func DecTopicUpvote(name string) bool {
	if v, ok := topicKV[name]; ok {
		if v.upvote != 0 {
			v.upvote = v.upvote - 1
		}
		return true
	}
	return false
}

// DecTopicDownvote decreases topic downvote counts
func DecTopicDownvote(name string) bool {
	if v, ok := topicKV[name]; ok {
		if v.downvote != 0 {
			v.downvote = v.downvote - 1
		}
		return true
	}
	return false
}
