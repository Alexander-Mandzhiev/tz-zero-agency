package news_repository

import (
	"context"
	"fmt"
	"strconv"
	"tz-zero-agency/internal/entity"
)

func (r *NewsRepository) GetAll(ctx context.Context, userID string, limit string, offset string) ([]entity.News, error) {
	const op = "repository.GetAll"
	var news []entity.News

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, fmt.Errorf("%s: invalid limit value: %w", op, err)
	}

	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return nil, fmt.Errorf("%s: invalid offset value: %w", op, err)
	}

	query := "SELECT * FROM news WHERE user_id = $1 ORDER BY id LIMIT $2 OFFSET $3"
	rows, err := r.db.Query(ctx, query, userID, limitInt, offsetInt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var n entity.News
		if err := rows.Scan(&n.ID, &n.UserID, &n.Title, &n.Content); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		n.Categories, err = r.GetCategoryIDs(ctx, int(n.ID))
		if err != nil {
			return nil, fmt.Errorf("%s: failed to get category IDs for news ID %d: %w", op, n.ID, err)
		}

		news = append(news, n)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return news, nil
}

func (r *NewsRepository) GetCategoryIDs(ctx context.Context, newsID int) ([]int, error) {
	const op = "repository.GetCategoryIDs"
	var categoryIDs []int

	query := "SELECT category_id FROM news_categories WHERE news_id = $1"
	rows, err := r.db.Query(ctx, query, newsID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var categoryID int
		if err := rows.Scan(&categoryID); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		categoryIDs = append(categoryIDs, categoryID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return categoryIDs, nil
}
