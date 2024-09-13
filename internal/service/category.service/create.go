package category_service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"tz-zero-agency/internal/entity"
	category_repository "tz-zero-agency/internal/repository/category.repository"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"
)

func (s *CategoryService) Create(ctx context.Context, category *entity.Category, token string) error {
	const op = "category.Create"
	log := s.logger.With(slog.String("op", op))
	log.Info("Create category")

	claims, err := jwt.ParseToken(token, s.secret)
	if err != nil {
		s.logger.Error("Ошибка распарсинга токена: %v", logger.Err(err))
		return fmt.Errorf("неверный токен: %w", err)
	}

	userID, ok := claims["uid"].(string)
	if !ok {
		return fmt.Errorf("не удалось получить userID из токена")
	}

	if err := s.repo.Create(ctx, category, userID); err != nil {
		if errors.Is(err, category_repository.ErrCategoryExists) {
			log.Warn("категория уже существует", logger.Err(err))
			return fmt.Errorf("%s: %w", op, category_repository.ErrCategoryExists)
		}
		log.Error("не удалось создать категорию", logger.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
