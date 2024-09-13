package user_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *UserRepository) Create(ctx context.Context, email string, passwordHash []byte) (string, error) {
	const op = "repository.SaveUser"
	r.log.Info("Signin user")
	var pgErr *pgconn.PgError
	var id string
	query := "INSERT INTO users (id, username, email, password_hash) VALUES ($1, $2, $3, $4) RETURNING id"

	if err := r.db.QueryRow(ctx, query, uuid.New().String(), "", email, passwordHash).Scan(&id); err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return "", fmt.Errorf("%s: %w", op, ErrUserExists)
		}
		return "", fmt.Errorf("%s: %W", op, err)
	}

	return id, nil
}
