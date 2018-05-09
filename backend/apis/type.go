package apis

import "github.com/google/uuid"

type topicInfo struct {
	UID      uuid.UUID `json:"uid"`
	Name     string    `json:"name"`
	Upvote   uint64    `json:"upvote"`
	Downvote uint64    `json:"downvote"`
}
