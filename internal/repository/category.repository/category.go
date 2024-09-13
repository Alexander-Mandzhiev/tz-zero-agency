package category_repository

import (
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrCreateCategory     = errors.New("ошибка создания категории")
	ErrGetCategory        = errors.New("ошибка при получении категорий")
	ErrInvalidCredentials = errors.New("неверный формат заголовка авторизации")
	ErrInternalServerErr  = errors.New("проблема на сервере")
	ErrUserNotFound       = errors.New("категория не найдена")
	ErrCategoryExists     = errors.New("категория уже существует")
)

type CategoryRepository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewCategoryRepository(log *slog.Logger, db *pgxpool.Pool) *CategoryRepository {
	return &CategoryRepository{log: log, db: db}
}
