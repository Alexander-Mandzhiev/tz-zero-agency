package news_service

import (
	"context"
	"fmt"
	"log/slog"
	"tz-zero-agency/internal/entity"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"
)

func (s *NewsService) Update(ctx context.Context, token string, news *entity.News, id int) (*entity.News, error) {
	const op = "news.Update"
	log := s.logger.With(slog.String("op", op))
	log.Info("Update news")

	claims, err := jwt.ParseToken(token, s.secret)
	if err != nil {
		s.logger.Error("Ошибка распарсинга токена: %v", logger.Err(err))
		return nil, fmt.Errorf("неверный токен: %w", err)
	}

	userID, ok := claims["uid"].(string)
	if !ok {
		s.logger.Error("Ошибка получения userID из токена")
		return nil, fmt.Errorf("не удалось получить userID из токена")
	}

	news.ID = id
	news.UserID = userID

	n, err := s.repo.Update(ctx, news)
	if err != nil {
		log.Error("не удалось обновить новость", logger.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return n, nil
}
