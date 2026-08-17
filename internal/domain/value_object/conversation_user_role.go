package value_object

import (
	"errors"
	"strings"
)

type ConversationUserRole struct {
	value string
}

const (
	conversationUserRoleOwner  = "owner"
	conversationUserRoleAdmin  = "admin"
	conversationUserRoleMember = "member"
)

var (
	ConversationUserRoleOwner = ConversationUserRole{
		value: conversationUserRoleOwner,
	}
	ConversationUserRoleAdmin = ConversationUserRole{
		value: conversationUserRoleAdmin,
	}
	ConversationUserRoleMember = ConversationUserRole{
		value: conversationUserRoleMember,
	}
)

func NewConversationUserRole(value string) (ConversationUserRole, error) {
	value = strings.ToLower(strings.TrimSpace(value))

	switch value {
	case conversationUserRoleOwner:
		return ConversationUserRoleOwner, nil
	case conversationUserRoleAdmin:
		return ConversationUserRoleAdmin, nil
	case conversationUserRoleMember:
		return ConversationUserRoleMember, nil
	default:
		return ConversationUserRole{}, errors.New("invalid conversation member role")
	}
}

func (r ConversationUserRole) Value() string {
	return r.value
}

func (r ConversationUserRole) String() string {
	return r.value
}

func (r ConversationUserRole) IsOwner() bool {
	return r == ConversationUserRoleOwner
}

func (r ConversationUserRole) IsAdmin() bool {
	return r == ConversationUserRoleAdmin
}

func (r ConversationUserRole) IsMember() bool {
	return r == ConversationUserRoleMember
}
