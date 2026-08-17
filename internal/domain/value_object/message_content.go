package value_object

import (
	"errors"
	"strings"
)

type MessageContent struct {
	value string
}

func NewMessageContent(value string) (MessageContent, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return MessageContent{}, errors.New("message content cannot be empty")
	}

	return MessageContent{
		value: value,
	}, nil
}

func (m MessageContent) Value() string {
	return m.value
}

func (m MessageContent) String() string {
	return m.value
}
