package value_object

import "github.com/google/uuid"

type ConversationID struct {
	value uuid.UUID
}

func NewConversationID(value uuid.UUID) ConversationID {
	return ConversationID{
		value: value,
	}
}

func (id ConversationID) Value() uuid.UUID {
	return id.value
}

func (id ConversationID) String() string {
	return id.value.String()
}
