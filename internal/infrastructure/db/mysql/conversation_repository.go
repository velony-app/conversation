package mysql

import (
	"context"
	"database/sql"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationRepository struct {
	db *sql.DB
}

func NewConversationRepository(db *sql.DB) repository.ConversationRepository {
	return &ConversationRepository{db: db}
}

func (repo *ConversationRepository) Find(
	ctx context.Context,
	conversationID value_object.ConversationID,
) (*entity.Conversation, error) {
	return nil, nil
}

func (repo *ConversationRepository) Save(
	ctx context.Context,
	conversation *entity.Conversation,
) error {
	return nil
}
