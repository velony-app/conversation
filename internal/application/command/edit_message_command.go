package command

import "github.com/velony-app/conversation/internal/application/common"

type EditMessageCommand struct {
	ActorID string

	MessageID string
	Content   string
}

type EditMessageCommandResult struct {
	Message *common.MessageResult
}
