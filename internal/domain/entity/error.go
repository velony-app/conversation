package entity

import "errors"

var (
	ErrConversationDeleted = errors.New(
		"conversation is deleted",
	)
	ErrMessageDeleted = errors.New(
		"message is deleted",
	)
	ErrConversationUserAlreadyLeft = errors.New(
		"conversation member has already left the conversation",
	)
	ErrConversationUserNotInConversation = errors.New(
		"conversation member does not belong to this conversation",
	)
	ErrConversationUsersNotInSameConversation = errors.New(
		"conversation members do not belong to the same conversation",
	)
	ErrConversationUserCannotManageConversation = errors.New(
		"conversation member does not have permission to manage the conversation",
	)
	ErrConversationUserCannotDeleteConversation = errors.New(
		"conversation member does not have permission to delete the conversation",
	)
	ErrConversationUserCannotTransferOwnership = errors.New(
		"conversation member cannot transfer ownership",
	)
	ErrConversationOwnerCannotLeave = errors.New(
		"conversation owner cannot leave the conversation",
	)
	ErrConversationUserCannotKick = errors.New(
		"conversation member does not have permission to kick members",
	)
	ErrConversationUserCannotKickSelf = errors.New(
		"conversation member cannot kick itself",
	)
	ErrConversationUserCannotKickOwner = errors.New(
		"conversation owner cannot be kicked",
	)
	ErrConversationUserCannotEditMessage = errors.New(
		"conversation member does not have permission to edit this message",
	)
	ErrConversationUserCannotDeleteMessage = errors.New(
		"conversation member does not have permission to delete this message",
	)
	ErrMessageNotInConversation = errors.New(
		"message does not belong to this conversation",
	)
)
