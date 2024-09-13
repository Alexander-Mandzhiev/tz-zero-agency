package news_repository

import (
	"context"
	"errors"
	"fmt"
	"tz-zero-agency/internal/entity"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *NewsRepository) Create(ctx context.Context, news *entity.News) error {
	const op = "repository.news.Create"
	r.log.Info("Create news")
	var pgErr *pgconn.PgError
	var query string

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("%s: failed to begin transaction: %w", op, err)
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
			return
		}
		err = tx.Commit(ctx)
	}()

	for _, categoryID := range news.Categories {
		var count int
		query = "SELECT COUNT(*) FROM categories WHERE id = $1 AND user_id = $2"
		err = tx.QueryRow(ctx, query, categoryID, news.UserID).Scan(&count)
		if err != nil {
			return fmt.Errorf("%s: failed to check category existence: %w", op, err)
		}
		if count == 0 {
			return fmt.Errorf("%s: category with id %d does not exist for user %s", op, categoryID, news.UserID)
		}
	}

	query = "INSERT INTO news (user_id, title, content) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRow(ctx, query, news.UserID, news.Title, news.Content).Scan(&news.ID)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return fmt.Errorf("%s: %w", op, ErrExistNews)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	for _, categoryID := range news.Categories {
		query = "INSERT INTO news_categories (news_id, category_id) VALUES ($1, $2)"
		_, err = tx.Exec(ctx, query, news.ID, categoryID)
		if err != nil {
			return fmt.Errorf("%s: failed to insert into news_categories: %w", op, err)
		}
	}

	return nil
}
