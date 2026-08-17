package event

import "github.com/velony-app/conversation/internal/domain/value_object"

const (
	ConversationUserAggregateType = "conversation_user"

	ConversationUserAddedEventType   = "conversation_user.added"
	ConversationUserRemovedEventType = "conversation_user.removed"
)

type ConversationUserAddedEvent struct {
	BaseEvent

	Role value_object.ConversationUserRole
}

func NewConversationUserAddedEvent(
	conversationUserID value_object.ConversationUserID,
	role value_object.ConversationUserRole,
) ConversationUserAddedEvent {
	return ConversationUserAddedEvent{
		BaseEvent: NewBaseEvent(conversationUserID.String()),

		Role: role,
	}
}

func (e ConversationUserAddedEvent) Type() string {
	return ConversationUserAddedEventType
}

func (e ConversationUserAddedEvent) AggregateType() string {
	return ConversationUserAggregateType
}

type ConversationUserRemovedEvent struct {
	BaseEvent
}

func NewConversationUserRemovedEvent(
	conversationUserID value_object.ConversationUserID,
) ConversationUserRemovedEvent {
	return ConversationUserRemovedEvent{
		BaseEvent: NewBaseEvent(conversationUserID.String()),
	}
}

func (e ConversationUserRemovedEvent) Type() string {
	return ConversationUserRemovedEventType
}

func (e ConversationUserRemovedEvent) AggregateType() string {
	return ConversationUserAggregateType
}
