package common

import "time"

type MessageResult struct {
	ID             string
	ConversationID string
	UserID         string
	Content        string
	CreateTime     time.Time
	UpdateTime     time.Time
}
