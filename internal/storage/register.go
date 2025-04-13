package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (db *AvitoDB) Register(ctx context.Context, newUser *models.User) error {
	emailExists, err := db.checkEmail(ctx, newUser.Email)
	if err != nil {
		return err
	}
	if emailExists {
		return errors.New("пользователь уже зарегистрирован")
	}

	_, err = db.pool.Exec(ctx, `
			INSERT INTO users (id, email, password, role)
			VALUES ($1, $2, $3, $4)
	`, newUser.ID, newUser.Email, newUser.Password, newUser.Role)

	return err
}

func (db *AvitoDB) checkEmail(ctx context.Context, email string) (bool, error) {
	var emailUser string
	err := db.pool.QueryRow(ctx, `
			SELECT email
			FROM users
			WHERE email = $1
			LIMIT 1
	`, email).Scan(&emailUser)

	if err == nil {
		return true, nil // такой пользователь уже есть
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil // такого пользователя нет
	}

	// возврат реальной ошибки
	return false, err
}
