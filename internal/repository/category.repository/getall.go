package category_repository

import (
	"context"
	"fmt"
	"tz-zero-agency/internal/entity"
)

func (r *CategoryRepository) GetAll(ctx context.Context, userId string) ([]entity.Category, error) {
	const op = "repository.category.GetAll"
	r.log.Info("GetAll category")

	var categories []entity.Category

	query := "SELECT id, user_id, title FROM categories WHERE user_id = $1 ORDER BY id"
	rows, err := r.db.Query(ctx, query, userId)
	if err != nil {
		r.log.Error(fmt.Sprintf("%s: failed to execute query: %v", op, err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var c entity.Category
		if err := rows.Scan(&c.ID, &c.UserID, &c.Title); err != nil {
			r.log.Error(fmt.Sprintf("%s: failed to scan row: %v", op, err))
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		categories = append(categories, c)
	}
	if err := rows.Err(); err != nil {
		r.log.Error(fmt.Sprintf("%s: error while iterating over rows: %v", op, err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return categories, nil
}
