package command

type RemoveConversationUserCommand struct {
	ActorID string

	ConversationID string
	UserID         string
}

type RemoveConversationUserCommandResult struct {
}
