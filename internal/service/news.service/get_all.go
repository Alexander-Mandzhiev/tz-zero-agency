package news_service

import (
	"context"
	"fmt"
	"tz-zero-agency/internal/entity"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"
)

func (s *NewsService) GetAll(ctx context.Context, token, limit, offset string) ([]entity.News, error) {
	var news []entity.News
	claims, err := jwt.ParseToken(token, s.secret)
	if err != nil {
		s.logger.Error("Ошибка распарсинга токена: %v", logger.Err(err))
		return nil, fmt.Errorf("неверный токен: %w", err)
	}
	userID := claims["uid"].(string)

	news, err = s.repo.GetAll(ctx, userID, limit, offset)
	if err != nil {
		return news, fmt.Errorf("неверный токен: %w", err)
	}
	return news, nil
}
