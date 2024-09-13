package user_repository

import (
	"context"
	"fmt"
	"tz-zero-agency/internal/entity"

	"github.com/jackc/pgx/v5"
)

func (r *UserRepository) User(ctx context.Context, email string) (entity.User, error) {
	const op = "repository.User"
	query := `SELECT id, username, email, password_hash FROM users WHERE email = $1`
	var u entity.User
	if err := r.db.QueryRow(context.Background(), query, email).Scan(&u.ID, &u.Username, &u.Email, &u.PasswordHash); err != nil {
		if err == pgx.ErrNoRows {
			return entity.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}
		return entity.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return u, nil
}
