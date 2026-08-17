package repository

import (
	"context"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationUserRepository interface {
	Find(ctx context.Context, conversationUserID value_object.ConversationUserID) (*entity.ConversationUser, error)
	FindByConversationAndUser(ctx context.Context, conversationID value_object.ConversationID, userID value_object.UserID) (*entity.ConversationUser, error)
	Save(ctx context.Context, conversationUser *entity.ConversationUser) error
}
