package service

import (
	"context"
	"log/slog"
	"time"
	"tz-zero-agency/internal/entity"
	"tz-zero-agency/internal/repository"
	category_service "tz-zero-agency/internal/service/category.service"
	news_service "tz-zero-agency/internal/service/news.service"
	user_service "tz-zero-agency/internal/service/user.service"
)

type User interface {
	Signup(ctx context.Context, email, password string) (string, error)
	Signin(ctx context.Context, email, password string) (string, error)
}

type News interface {
	Create(ctx context.Context, news *entity.News, token string) error
	GetAll(ctx context.Context, token, limit, offset string) ([]entity.News, error)
	Update(ctx context.Context, token string, news *entity.News, id int) (*entity.News, error)
}

type Categoryes interface {
	Create(ctx context.Context, category *entity.Category, token string) error
	GetAll(ctx context.Context, token string) ([]entity.Category, error)
}

type Service struct {
	User
	News
	Categoryes
}

func NewService(repository *repository.Repository, log *slog.Logger, tokenTTL time.Duration, secret string) *Service {
	return &Service{
		User:       user_service.NewUserService(repository.User, log, tokenTTL, secret),
		News:       news_service.NewNewsService(repository.News, log, tokenTTL, secret),
		Categoryes: category_service.NewCategoryService(repository.Categoryes, log, tokenTTL, secret),
	}
}
