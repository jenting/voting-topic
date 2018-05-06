package apis

type topicInfo struct {
	Name     string `json:"name" binding:"required"`
	Upvote   uint64 `json:"upvote"`
	Downvote uint64 `json:"downvote"`
}
