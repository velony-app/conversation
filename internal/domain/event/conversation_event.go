package event

import (
	"time"

	"github.com/velony-app/conversation/internal/domain/value_object"
)

const (
	ConversationAggregateType = "conversation"

	ConversationCreatedEventType            = "conversation.created"
	ConversationTitleChangedEventType       = "conversation.title.changed"
	ConversationAvatarImageChangedEventType = "conversation.avatar_image.changed"
	ConversationDeletedEventType            = "conversation.deleted"
)

type ConversationCreatedEvent struct {
	BaseEvent

	Title       value_object.ConversationTitle
	AvatarImage value_object.ResourceName
	CreateTime  time.Time
}

func NewConversationCreatedEvent(
	conversationID value_object.ConversationID,
	title value_object.ConversationTitle,
	avatarImage value_object.ResourceName,
	createTime time.Time,
) ConversationCreatedEvent {
	return ConversationCreatedEvent{
		BaseEvent: NewBaseEvent(conversationID.String()),

		Title:       title,
		AvatarImage: avatarImage,
		CreateTime:  createTime,
	}
}

func (e ConversationCreatedEvent) Type() string {
	return ConversationCreatedEventType
}

func (e ConversationCreatedEvent) AggregateType() string {
	return ConversationAggregateType
}

type ConversationTitleChangedEvent struct {
	BaseEvent

	Title      value_object.ConversationTitle
	UpdateTime time.Time
}

func NewConversationTitleChangedEvent(
	conversationID value_object.ConversationID,
	title value_object.ConversationTitle,
	updateTime time.Time,
) ConversationTitleChangedEvent {
	return ConversationTitleChangedEvent{
		BaseEvent: NewBaseEvent(conversationID.String()),

		Title:      title,
		UpdateTime: updateTime,
	}
}

func (e ConversationTitleChangedEvent) Type() string {
	return ConversationTitleChangedEventType
}

func (e ConversationTitleChangedEvent) AggregateType() string {
	return ConversationAggregateType
}

type ConversationAvatarImageChangedEvent struct {
	BaseEvent

	AvatarImage value_object.ResourceName
	UpdateTime  time.Time
}

func NewConversationAvatarImageChangedEvent(
	conversationID value_object.ConversationID,
	avatarImage value_object.ResourceName,
	updateTime time.Time,
) ConversationAvatarImageChangedEvent {
	return ConversationAvatarImageChangedEvent{
		BaseEvent: NewBaseEvent(conversationID.String()),

		AvatarImage: avatarImage,
		UpdateTime:  updateTime,
	}
}

func (e ConversationAvatarImageChangedEvent) Type() string {
	return ConversationAvatarImageChangedEventType
}

func (e ConversationAvatarImageChangedEvent) AggregateType() string {
	return ConversationAggregateType
}

type ConversationDeletedEvent struct {
	BaseEvent

	DeleteTime time.Time
}

func NewConversationDeletedEvent(
	conversationID value_object.ConversationID,
	deleteTime time.Time,
) ConversationDeletedEvent {
	return ConversationDeletedEvent{
		BaseEvent: NewBaseEvent(conversationID.String()),

		DeleteTime: deleteTime,
	}
}

func (e ConversationDeletedEvent) Type() string {
	return ConversationDeletedEventType
}

func (e ConversationDeletedEvent) AggregateType() string {
	return ConversationAggregateType
}
