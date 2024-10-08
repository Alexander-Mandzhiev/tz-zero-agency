package news_repository

import (
	"context"
	"errors"
	"fmt"
	"tz-zero-agency/internal/entity"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *NewsRepository) Update(ctx context.Context, news *entity.News) (*entity.News, error) {
	const op = "repository.news.Update"
	r.log.Info("Update news")
	var pgErr *pgconn.PgError
	var query string

	tx, err := r.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to begin transaction: %w", op, err)
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
			return nil, fmt.Errorf("%s: failed to check category existence: %w", op, err)
		}
		if count == 0 {
			return nil, fmt.Errorf("%s: category with id %d does not exist for user %s", op, categoryID, news.UserID)
		}
	}

	query = "UPDATE news SET user_id = $1, title = $2, content = $3 WHERE id = $4 RETURNING id"
	err = tx.QueryRow(ctx, query, news.UserID, news.Title, news.Content, news.ID).Scan(&news.ID)
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return nil, fmt.Errorf("%s: %w", op, ErrExistNews)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Формируем запрос для обновления категорий
	values := make([]string, 0, len(news.Categories))
	args := make([]interface{}, 0, len(news.Categories)*2)
	for i, categoryID := range news.Categories {
		values = append(values, fmt.Sprintf("($%d, $%d)", i*2+1, i*2+2))
		args = append(args, news.ID, categoryID)
	}

	query = fmt.Sprintf("INSERT INTO news_categories (news_id, category_id) VALUES %s ON CONFLICT (news_id, category_id) DO UPDATE SET category_id = EXCLUDED.category_id",
		join(values, ","))
	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to upsert into news_categories: %w", op, err)
	}

	return news, nil
}

func join(elements []string, separator string) string {
	if len(elements) == 0 {
		return ""
	}
	if len(elements) == 1 {
		return elements[0]
	}
	result := elements[0]
	for _, element := range elements[1:] {
		result += separator + element
	}
	return result
}
