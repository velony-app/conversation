package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/velony-app/conversation/internal/domain/event"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type Message struct {
	ID             value_object.MessageID
	ConversationID value_object.ConversationID
	UserID         value_object.UserID
	Content        value_object.MessageContent
	CreateTime     time.Time
	UpdateTime     time.Time
	DeleteTime     *time.Time

	domainEvents []event.DomainEvent
}

func NewMessage(actor *ConversationUser, content value_object.MessageContent) (*Message, error) {
	if actor.LeaveTime != nil {
		return nil, ErrConversationUserAlreadyLeft
	}

	now := time.Now()
	messageID := value_object.NewMessageID(uuid.Must(uuid.NewV7()))

	message := &Message{
		ID:             messageID,
		ConversationID: actor.ID.ConversationID(),
		UserID:         actor.ID.UserID(),
		Content:        content,
		CreateTime:     now,
		UpdateTime:     now,
	}

	message.recordEvent(
		event.NewMessageSentEvent(
			messageID,
			actor.ID.ConversationID(),
			actor.ID.UserID(),
			content,
			now,
		),
	)

	return message, nil
}

func (m *Message) Edit(actor *ConversationUser, content value_object.MessageContent) error {
	if m.DeleteTime != nil {
		return ErrMessageDeleted
	}
	if actor.LeaveTime != nil {
		return ErrConversationUserAlreadyLeft
	}
	if actor.ID.ConversationID() != m.ConversationID {
		return ErrConversationUserNotInConversation
	}
	if actor.ID.UserID() != m.UserID {
		return ErrConversationUserCannotEditMessage
	}

	now := time.Now()

	m.Content = content
	m.UpdateTime = now

	m.recordEvent(
		event.NewMessageEditedEvent(
			m.ID,
			content,
			now,
		),
	)

	return nil
}

func (m *Message) Delete(actor *ConversationUser) error {
	if m.DeleteTime != nil {
		return ErrMessageDeleted
	}
	if actor.LeaveTime != nil {
		return ErrConversationUserAlreadyLeft
	}
	if actor.ID.ConversationID() != m.ConversationID {
		return ErrConversationUserNotInConversation
	}
	if actor.ID.UserID() != m.UserID && !actor.Role.IsAdmin() && !actor.Role.IsOwner() {
		return ErrConversationUserCannotDeleteMessage
	}

	now := time.Now()

	m.UpdateTime = now
	m.DeleteTime = &now

	m.recordEvent(
		event.NewMessageDeletedEvent(
			m.ID,
			now,
		),
	)

	return nil
}

func (m *Message) PullEvents() []event.DomainEvent {
	pulled := m.domainEvents
	m.domainEvents = nil
	return pulled
}

func (m *Message) recordEvent(domainEvent event.DomainEvent) {
	m.domainEvents = append(m.domainEvents, domainEvent)
}
