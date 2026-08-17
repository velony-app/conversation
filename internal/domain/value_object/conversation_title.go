package value_object

import (
	"errors"
	"strings"
)

type ConversationTitle struct {
	value string
}

func NewConversationTitle(value string) (ConversationTitle, error) {
	value = strings.TrimSpace(value)

	if value == "" {
		return ConversationTitle{}, errors.New("conversation title cannot be empty")
	}

	if len(value) > 255 {
		return ConversationTitle{}, errors.New("conversation title cannot exceed 255 characters")
	}

	return ConversationTitle{
		value: value,
	}, nil
}

func (t ConversationTitle) Value() string {
	return t.value
}

func (t ConversationTitle) String() string {
	return t.value
}
