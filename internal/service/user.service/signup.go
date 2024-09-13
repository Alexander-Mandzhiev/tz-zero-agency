package user_service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	user_repository "tz-zero-agency/internal/repository/user.repository"
	"tz-zero-agency/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Signup(ctx context.Context, email, password string) (string, error) {
	const op = "auth.Signup"
	log := s.logger.With(slog.String("op", op))
	log.Info("Signup user")

	if err := validate_data(email, password); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("не удалось сгенерировать хэш пароля", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.repo.Create(ctx, email, passwordHash)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserExists) {
			log.Warn("пользователь уже существует", logger.Err(err))
			return "", fmt.Errorf("%s: %w", op, user_repository.ErrUserExists)
		}
		log.Error("не удалось создать пользователя", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
