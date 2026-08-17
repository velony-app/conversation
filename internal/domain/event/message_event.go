package event

import (
	"time"

	"github.com/velony-app/conversation/internal/domain/value_object"
)

const (
	MessageAggregateType = "message"

	MessageSentEventType    = "message.sent"
	MessageEditedEventType  = "message.edited"
	MessageDeletedEventType = "message.deleted"
)

type MessageSentEvent struct {
	BaseEvent

	ConversationID value_object.ConversationID
	UserID         value_object.UserID
	Content        value_object.MessageContent
	CreateTime     time.Time
}

func NewMessageSentEvent(
	messageID value_object.MessageID,
	conversationID value_object.ConversationID,
	userID value_object.UserID,
	content value_object.MessageContent,
	createTime time.Time,
) MessageSentEvent {
	return MessageSentEvent{
		BaseEvent: NewBaseEvent(messageID.String()),

		ConversationID: conversationID,
		UserID:         userID,
		Content:        content,
		CreateTime:     createTime,
	}
}

func (e MessageSentEvent) Type() string {
	return MessageSentEventType
}

func (e MessageSentEvent) AggregateType() string {
	return MessageAggregateType
}

type MessageEditedEvent struct {
	BaseEvent

	Content    value_object.MessageContent
	UpdateTime time.Time
}

func NewMessageEditedEvent(
	messageID value_object.MessageID,
	content value_object.MessageContent,
	updateTime time.Time,
) MessageEditedEvent {
	return MessageEditedEvent{
		BaseEvent: NewBaseEvent(messageID.String()),

		Content:    content,
		UpdateTime: updateTime,
	}
}

func (e MessageEditedEvent) Type() string {
	return MessageEditedEventType
}

func (e MessageEditedEvent) AggregateType() string {
	return MessageAggregateType
}

type MessageDeletedEvent struct {
	BaseEvent

	DeleteTime time.Time
}

func NewMessageDeletedEvent(
	messageID value_object.MessageID,
	deleteTime time.Time,
) MessageDeletedEvent {
	return MessageDeletedEvent{
		BaseEvent: NewBaseEvent(messageID.String()),

		DeleteTime: deleteTime,
	}
}

func (e MessageDeletedEvent) Type() string {
	return MessageDeletedEventType
}

func (e MessageDeletedEvent) AggregateType() string {
	return MessageAggregateType
}
