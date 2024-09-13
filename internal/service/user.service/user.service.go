package user_service

import (
	"errors"
	"log/slog"
	"time"
	"tz-zero-agency/internal/repository"
	"tz-zero-agency/pkg/validate"
)

type UserService struct {
	logger   *slog.Logger
	repo     repository.User
	tokenTTL time.Duration
	secret   string
}

func NewUserService(repo repository.User, logger *slog.Logger, tokenTTL time.Duration, secret string) *UserService {
	return &UserService{logger: logger, repo: repo, tokenTTL: tokenTTL, secret: secret}
}

func validate_data(email, password string) error {
	if !validate.IsValidEmail(email) {
		return errors.New("invalid email")
	}
	if !validate.IsValidPassword(password) {
		return errors.New("the password must be longer than 7 characters and contain lowercase letters, uppercase letters, numbers and special characters")
	}
	return nil
}
