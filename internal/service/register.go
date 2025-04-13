package service

import (
	"avitopvz/internal/auth"
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

func (s *service) Register(ctx context.Context, newUser *models.User) (string, error) {
	idUUID, _ := uuid.NewV4()
	newUser.ID = idUUID

	if newUser.Role != "moderator" && newUser.Role != "employee" {
		return "", errors.New("нет такой роли")
	}

	hashedPassword, err := auth.HashPassword(newUser.Password)
	if err != nil {
		return "", err
	}
	newUser.Password = hashedPassword

	err = s.storage.Register(ctx, newUser)
	if err != nil {
		return "", err
	}

	key, err := auth.GetJWTKey()
	if err != nil {
		return "", err
	}

	tokenString, err := auth.CreateToken(newUser, key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
