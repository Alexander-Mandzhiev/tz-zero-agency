package validate

import (
	"fmt"
	"net/mail"
	"tz-zero-agency/internal/entity"
	"unicode"
)

func IsValidPassword(plaintext string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	s := []rune(plaintext)
	if len(s) >= 7 {
		hasMinLen = true
	}
	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidateNews(news *entity.News) error {
	if news.Title == "" {
		return fmt.Errorf("заголовок не может быть пустым")
	}
	if news.Content == "" {
		return fmt.Errorf("контент не может быть пустым")
	}
	if len(news.Title) > 255 {
		return fmt.Errorf("заголовок не должен превышать 255 символов")
	}
	if len(news.Content) > 2000 {
		return fmt.Errorf("контент не должен превышать 2000 символов")
	}
	return nil
}

func ValidateCategories(news *entity.Category) error {
	if news.Title == "" {
		return fmt.Errorf("заголовок не может быть пустым")
	}

	if len(news.Title) > 255 {
		return fmt.Errorf("заголовок не должен превышать 255 символов")
	}

	return nil
}
