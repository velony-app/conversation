package repository

import (
	"context"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type MessageRepository interface {
	Find(ctx context.Context, message value_object.MessageID) (*entity.Message, error)
	Save(ctx context.Context, message *entity.Message) error
}
