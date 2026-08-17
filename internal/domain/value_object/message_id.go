package value_object

import "github.com/google/uuid"

type MessageID struct {
	value uuid.UUID
}

func NewMessageID(value uuid.UUID) MessageID {
	return MessageID{
		value: value,
	}
}

func (id MessageID) Value() uuid.UUID {
	return id.value
}

func (id MessageID) String() string {
	return id.value.String()
}
