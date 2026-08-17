package entity

import (
	"time"

	"github.com/google/uuid"

	"github.com/velony-app/conversation/internal/domain/event"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type Conversation struct {
	ID          value_object.ConversationID
	Title       value_object.ConversationTitle
	AvatarImage value_object.ResourceName
	CreateTime  time.Time
	UpdateTime  time.Time
	DeleteTime  *time.Time

	domainEvents []event.DomainEvent
}

func NewConversation(
	ownerID value_object.UserID,
	title value_object.ConversationTitle,
	avatarImage value_object.ResourceName,
) *Conversation {
	now := time.Now()
	conversationID := value_object.NewConversationID(uuid.Must(uuid.NewV7()))

	conversation := &Conversation{
		ID:          conversationID,
		Title:       title,
		AvatarImage: avatarImage,
		CreateTime:  now,
		UpdateTime:  now,
	}

	conversation.recordEvent(
		event.NewConversationCreatedEvent(
			conversationID,
			title,
			avatarImage,
			now,
		),
	)

	return conversation
}

func (c *Conversation) ChangeTitle(conversationUser *ConversationUser, title value_object.ConversationTitle) error {
	if c.DeleteTime != nil {
		return ErrConversationDeleted
	}
	if conversationUser.LeaveTime != nil {
		return ErrConversationUserAlreadyLeft
	}
	if conversationUser.ID.ConversationID() != c.ID {
		return ErrConversationUserNotInConversation
	}
	if !conversationUser.Role.IsAdmin() && !conversationUser.Role.IsOwner() {
		return ErrConversationUserCannotManageConversation
	}

	now := time.Now()

	c.Title = title
	c.UpdateTime = now

	c.recordEvent(
		event.NewConversationTitleChangedEvent(
			c.ID,
			title,
			now,
		),
	)

	return nil
}

func (c *Conversation) Delete(
	conversationUser *ConversationUser,
) error {
	if c.DeleteTime != nil {
		return ErrConversationDeleted
	}
	if conversationUser.LeaveTime != nil {
		return ErrConversationUserAlreadyLeft
	}
	if conversationUser.ID.ConversationID() != c.ID {
		return ErrConversationUserNotInConversation
	}
	if !conversationUser.Role.IsOwner() {
		return ErrConversationUserCannotDeleteConversation
	}

	now := time.Now()

	c.UpdateTime = now
	c.DeleteTime = &now

	c.recordEvent(
		event.NewConversationDeletedEvent(
			c.ID,
			now,
		),
	)

	return nil
}

func (c *Conversation) PullEvents() []event.DomainEvent {
	pulled := c.domainEvents
	c.domainEvents = nil
	return pulled
}

func (c *Conversation) recordEvent(domainEvent event.DomainEvent) {
	c.domainEvents = append(c.domainEvents, domainEvent)
}
