package news_service

import (
	"log/slog"
	"time"
	"tz-zero-agency/internal/repository"
)

type NewsService struct {
	logger   *slog.Logger
	repo     repository.News
	tokenTTL time.Duration
	secret   string
}

func NewNewsService(repo repository.News, logger *slog.Logger, tokenTTL time.Duration, secret string) *NewsService {
	return &NewsService{logger: logger, repo: repo, tokenTTL: tokenTTL, secret: secret}
}
