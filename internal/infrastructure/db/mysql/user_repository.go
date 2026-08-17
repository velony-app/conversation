package mysql

import (
	"context"
	"database/sql"
	"errors"

	"github.com/velony-app/conversation/internal/domain/entity"
	"github.com/velony-app/conversation/internal/domain/repository"
	"github.com/velony-app/conversation/internal/domain/value_object"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (repo *UserRepository) Find(ctx context.Context, userID value_object.UserID) (*entity.User, error) {
	const query = `
		SELECT id
		FROM users
		WHERE id = ?
		LIMIT 1
	`

	var user entity.User

	if err := executor(ctx, repo.db).QueryRowContext(ctx, query, userID.Value().String()).Scan(&user.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}
