package command

type DeleteConversationCommand struct {
	ActorID string

	ConversationID string
}

type DeleteConversationCommandResult struct {
}
