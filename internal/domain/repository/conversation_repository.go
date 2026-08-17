package repository

import (
	"context"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationRepository interface {
	Find(ctx context.Context, conversationID value_object.ConversationID) (*entity.Conversation, error)
	Save(ctx context.Context, conversation *entity.Conversation) error
}
