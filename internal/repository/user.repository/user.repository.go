package user_repository

import (
	"errors"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrInvalidCredentials = errors.New("неправильный формат данных")
	ErrInternalServerErr  = errors.New("проблема на сервере")
	ErrUserNotFound       = errors.New("пользователь не найдена")
	ErrUserExists         = errors.New("пользователь уже существует")
)

type UserRepository struct {
	log *slog.Logger
	db  *pgxpool.Pool
}

func NewUserRepository(log *slog.Logger, db *pgxpool.Pool) *UserRepository {
	return &UserRepository{log: log, db: db}
}
