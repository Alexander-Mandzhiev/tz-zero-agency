package user_service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	user_repository "tz-zero-agency/internal/repository/user.repository"
	"tz-zero-agency/pkg/jwt"
	"tz-zero-agency/pkg/logger"

	"golang.org/x/crypto/bcrypt"
)

func (s *UserService) Signin(ctx context.Context, email, password string) (string, error) {
	const op = "auth.Signin"
	log := s.logger.With(slog.String("op", op))
	log.Info("Signin user")

	if err := validate_data(email, password); err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	u, err := s.repo.User(ctx, email)
	if err != nil {
		if errors.Is(err, user_repository.ErrUserNotFound) {
			s.logger.Warn("пользователь не найден", logger.Err(err))
			return "", fmt.Errorf("%s: %w", op, user_repository.ErrInvalidCredentials)
		}
		s.logger.Warn("не удалось получить пользователя", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}

	if err := bcrypt.CompareHashAndPassword(u.PasswordHash, []byte(password)); err != nil {
		s.logger.Warn("invalid credentials", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, user_repository.ErrInvalidCredentials)
	}

	log.Info("Signin user is successfully")

	token, err := jwt.NewToken(u, s.secret, s.tokenTTL)
	if err != nil {
		s.logger.Warn("failed to generate token", logger.Err(err))
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return token, nil

}
