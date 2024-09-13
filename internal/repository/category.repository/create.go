package category_repository

import (
	"context"
	"errors"
	"fmt"
	"tz-zero-agency/internal/entity"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoryRepository) Create(ctx context.Context, category *entity.Category, userId string) error {
	const op = "repository.category.Create"
	r.log.Info("Create category")
	var pgErr *pgconn.PgError

	query := "INSERT INTO categories (user_id, title) VALUES ($1, $2) RETURNING *"
	err := r.db.QueryRow(ctx, query, userId, category.Title).Scan(&category.ID, &category.UserID, &category.Title)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%s: %w", op, ErrCategoryExists)
		}
		return fmt.Errorf("%s: %W", op, err)
	}

	return nil
}
