package query

import "github.com/velony-app/conversation/internal/application/common"

type GetConversationQuery struct {
	ActorID string

	ConversationID string
}

type GetConversationQueryResult struct {
	Conversation *common.ConversationResult
}
