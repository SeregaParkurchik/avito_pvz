package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"
)

func (db *AvitoDB) Login(ctx context.Context, user *models.User) (*models.User, error) {
	emailExists, err := db.checkEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if !emailExists {
		return nil, errors.New("пользователя с таким email не существует")
	}

	foundUser := &models.User{}
	err = db.pool.QueryRow(ctx, `
			SELECT id, email, password, role
			FROM users
			WHERE email = $1
	`, user.Email).Scan(&foundUser.ID, &foundUser.Email, &foundUser.Password, &foundUser.Role)

	if err != nil {
		return nil, err
	}

	return foundUser, nil
}
