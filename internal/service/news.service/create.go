package news_service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"tz-zero-agency/internal/entity"
	news_repository "tz-zero-agency/internal/repository/news.repository"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"
)

func (s *NewsService) Create(ctx context.Context, news *entity.News, token string) error {
	const op = "news.Create"
	log := s.logger.With(slog.String("op", op))
	log.Info("Create news")

	claims, err := jwt.ParseToken(token, s.secret)
	if err != nil {
		s.logger.Error("Ошибка распарсинга токена: %v", logger.Err(err))
		return fmt.Errorf("неверный токен: %w", err)
	}

	userID, ok := claims["uid"].(string)
	if !ok {
		return fmt.Errorf("не удалось получить userID из токена")
	}

	news.UserID = userID

	if err := s.repo.Create(ctx, news); err != nil {
		if errors.Is(err, news_repository.ErrExistNews) {
			log.Warn("новость уже создана", logger.Err(err))
			return fmt.Errorf("%s: %w", op, news_repository.ErrExistNews)
		}
		log.Error("не удалось создать новость", logger.Err(err))
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
