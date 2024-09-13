package repository

import (
	"context"
	"log/slog"
	"tz-zero-agency/internal/entity"
	category_repository "tz-zero-agency/internal/repository/category.repository"
	news_repository "tz-zero-agency/internal/repository/news.repository"
	user_repository "tz-zero-agency/internal/repository/user.repository"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User interface {
	Create(ctx context.Context, email string, passwordHash []byte) (string, error)
	User(ctx context.Context, email string) (entity.User, error)
}

type News interface {
	Create(ctx context.Context, news *entity.News) error
	GetAll(ctx context.Context, user_id, limit, offset string) ([]entity.News, error)
	Update(ctx context.Context, news *entity.News) (*entity.News, error)
}

type Categoryes interface {
	Create(ctx context.Context, category *entity.Category, userId string) error
	GetAll(ctx context.Context, userId string) ([]entity.Category, error)
}

type Repository struct {
	User
	News
	Categoryes
}

func NewRepository(log *slog.Logger, db *pgxpool.Pool) *Repository {
	return &Repository{
		User:       user_repository.NewUserRepository(log, db),
		News:       news_repository.NewNewsRepository(log, db),
		Categoryes: category_repository.NewCategoryRepository(log, db),
	}
}
