package service

import (
	"avitopvz/internal/auth"
	"avitopvz/internal/models"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, user *models.User) (string, error) {
	foundUser, err := s.storage.Login(ctx, user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		return "", errors.New("неверный пароль")
	}

	key, err := auth.GetJWTKey()
	if err != nil {
		return "", err
	}

	tokenString, err := auth.CreateToken(foundUser, key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
