package command

import "github.com/velony-app/conversation/internal/application/common"

type CreateConversationCommand struct {
	ActorID string

	Title  string
	Avatar string
}

type CreateConversationCommandResult struct {
	Conversation *common.ConversationResult
}
