package news_repository

import (
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrCreateNews         = errors.New("ошибка создания новости")
	ErrGetNews            = errors.New("ошибка при получении новостей")
	ErrInvalidCredentials = errors.New("неверный формат заголовка авторизации")
	ErrInternalServerErr  = errors.New("проблема на сервере")
	ErrUserNotFound       = errors.New("пользователь не найдена")
	ErrExistNews          = errors.New("новость уже новости")
)

type NewsRepository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewNewsRepository(log *slog.Logger, db *pgxpool.Pool) *NewsRepository {
	return &NewsRepository{log: log, db: db}
}
