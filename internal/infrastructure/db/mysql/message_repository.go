package mysql

import (
	"context"
	"database/sql"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type MessageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) repository.MessageRepository {
	return &MessageRepository{db: db}
}

func (repo *MessageRepository) Find(
	ctx context.Context,
	messageID value_object.MessageID,
) (*entity.Message, error) {
	return nil, nil
}

func (repo *MessageRepository) Save(
	ctx context.Context,
	message *entity.Message,
) error {
	return nil
}
