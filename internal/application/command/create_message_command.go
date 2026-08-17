package command

import "github.com/velony-app/conversation/internal/application/common"

type SendMessageCommand struct {
	ActorID string

	ConversationID string
	Content        string
}

type SendMessageCommandResult struct {
	Message *common.MessageResult
}
