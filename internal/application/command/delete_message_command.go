package command

type DeleteMessageCommand struct {
	ActorID string

	MessageID string
}

type DeleteMessageCommandResult struct {
}
