package mysql

import (
	"context"
	"database/sql"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type ConversationUserRepository struct {
	db *sql.DB
}

func NewConversationUserRepository(
	db *sql.DB,
) repository.ConversationUserRepository {
	return &ConversationUserRepository{
		db: db,
	}
}

func (repo *ConversationUserRepository) Find(
	ctx context.Context,
	conversationUserID value_object.ConversationUserID,
) (*entity.ConversationUser, error) {
	return nil, nil
}

func (repo *ConversationUserRepository) FindByConversationAndUser(
	ctx context.Context,
	conversationID value_object.ConversationID,
	userID value_object.UserID,
) (*entity.ConversationUser, error) {
	return nil, nil
}

func (repo *ConversationUserRepository) Save(
	ctx context.Context,
	conversationUser *entity.ConversationUser,
) error {
	return nil
}
