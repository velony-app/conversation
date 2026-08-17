package common

import "time"

type ConversationUserResult struct {
	ID       string
	Role     string
	JoinTime time.Time
}
