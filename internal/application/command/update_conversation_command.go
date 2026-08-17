package command

import "github.com/velony-app/conversation/internal/application/common"

type UpdateConversationCommand struct {
	ActorID string

	ConversationID string
	Title          *string
}

type UpdateConversationCommandResult struct {
	Conversation *common.ConversationResult
}
