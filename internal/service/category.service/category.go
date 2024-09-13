package category_service

import (
	"log/slog"
	"time"
	"tz-zero-agency/internal/repository"
)

type CategoryService struct {
	logger   *slog.Logger
	repo     repository.Categoryes
	tokenTTL time.Duration
	secret   string
}

func NewCategoryService(repo repository.Categoryes, logger *slog.Logger, tokenTTL time.Duration, secret string) *CategoryService {
	return &CategoryService{logger: logger, repo: repo, tokenTTL: tokenTTL, secret: secret}
}
