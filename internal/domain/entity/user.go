package entity

import (
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type User struct {
	ID value_object.UserID
}

func (u *User) CreateConversation(title value_object.ConversationTitle, avatar value_object.ResourceName) (*Conversation, *ConversationUser) {
	conversation := NewConversation(
		u.ID,
		title,
		avatar,
	)
	conversationUser := NewConversationUser(
		value_object.NewConversationUserID(conversation.ID, u.ID),
		value_object.ConversationUserRoleOwner,
	)

	return conversation, conversationUser
}

func (u *User) JoinConversation(conversation *Conversation) (*ConversationUser, error) {
	if conversation.DeleteTime != nil {
		return nil, ErrConversationDeleted
	}

	conversationUser := NewConversationUser(
		value_object.NewConversationUserID(conversation.ID, u.ID),
		value_object.ConversationUserRoleOwner,
	)

	return conversationUser, nil
}
