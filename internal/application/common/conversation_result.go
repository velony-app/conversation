package common

import "time"

type ConversationResult struct {
	ID         string
	Title      string
	Avatar     string
	CreateTime time.Time
}
