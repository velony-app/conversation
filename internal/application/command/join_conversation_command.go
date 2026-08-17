package command

import "github.com/velony-app/conversation/internal/application/common"

type JoinConversationCommand struct {
	ActorID string

	ConversationID string
}

type JoinConversationCommandResult struct {
	ConversationUser *common.ConversationUserResult
}
