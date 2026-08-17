package command

type LeaveConversationCommand struct {
	ActorID string

	ConversationID string
}

type LeaveConversationCommandResult struct {
}
