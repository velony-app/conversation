package entity

import (
	"time"

	"github.com/velony-app/conversation/internal/domain/event"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationUser struct {
	ID        value_object.ConversationUserID
	Role      value_object.ConversationUserRole
	JoinTime  time.Time
	LeaveTime *time.Time

	domainEvents []event.DomainEvent
}

func NewConversationUser(
	conversationUserID value_object.ConversationUserID,
	role value_object.ConversationUserRole,
) *ConversationUser {
	now := time.Now()

	conversationUser := &ConversationUser{
		ID:       conversationUserID,
		Role:     role,
		JoinTime: now,
	}

	conversationUser.recordEvent(
		event.NewConversationUserAddedEvent(
			conversationUserID,
			role,
		),
	)

	return conversationUser
}

func (u *ConversationUser) Leave() error {
	if u.LeaveTime != nil {
		return ErrConversationUserAlreadyLeft
	}
	if u.Role.IsOwner() {
		return ErrConversationOwnerCannotLeave
	}

	now := time.Now()
	u.LeaveTime = &now

	u.recordEvent(
		event.NewConversationUserRemovedEvent(
			u.ID,
		),
	)

	return nil
}

func (u *ConversationUser) PullEvents() []event.DomainEvent {
	pulled := u.domainEvents
	u.domainEvents = nil
	return pulled
}

func (u *ConversationUser) recordEvent(domainEvent event.DomainEvent) {
	u.domainEvents = append(u.domainEvents, domainEvent)
}
