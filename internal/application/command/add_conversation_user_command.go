package command

import "github.com/velony-app/conversation/internal/application/common"

type AddConversationUserCommand struct {
	ActorID string

	ConversationID string
	UserID         string
}

type AddConversationUserCommandResult struct {
	ConversationUser *common.ConversationUserResult
}
