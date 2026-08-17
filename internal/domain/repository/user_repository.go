package repository

import (
	"context"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type UserRepository interface {
	Find(ctx context.Context, userID value_object.UserID) (*entity.User, error)
}
