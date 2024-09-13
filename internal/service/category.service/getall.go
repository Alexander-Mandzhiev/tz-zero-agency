package category_service

import (
	"context"
	"fmt"
	"tz-zero-agency/internal/entity"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"
)

func (s *CategoryService) GetAll(ctx context.Context, token string) ([]entity.Category, error) {
	var category []entity.Category
	claims, err := jwt.ParseToken(token, s.secret)
	if err != nil {
		s.logger.Error("Ошибка распарсинга токена: %v", logger.Err(err))
		return nil, fmt.Errorf("неверный токен: %w", err)
	}
	userID := claims["uid"].(string)

	category, err = s.repo.GetAll(ctx, userID)
	if err != nil {
		return category, err
	}
	return category, nil
}
