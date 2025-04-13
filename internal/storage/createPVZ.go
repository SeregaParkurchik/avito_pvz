package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *AvitoDB) CreatePVZ(ctx context.Context, newPVZ *models.PVZ) (*models.PVZ, error) {
	exists, err := db.checkPVZID(ctx, newPVZ.ID)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New("ПВЗ с таким ID уже существует")
	}

	query := `
			INSERT INTO pvz (id, registration_date, city)
			VALUES ($1, $2, $3)
			RETURNING id, registration_date, city
	`

	createdPVZ := &models.PVZ{}
	err = db.pool.QueryRow(ctx, query, newPVZ.ID, newPVZ.RegistrationDate, newPVZ.City).Scan(&createdPVZ.ID, &createdPVZ.RegistrationDate, &createdPVZ.City)

	if err != nil {
		return nil, err
	}

	return createdPVZ, nil
}

func (db *AvitoDB) checkPVZID(ctx context.Context, id uuid.UUID) (bool, error) {
	var pvzID uuid.UUID
	err := db.pool.QueryRow(ctx, `
			SELECT id
			FROM pvz
			WHERE id = $1
			LIMIT 1
	`, id).Scan(&pvzID)

	if err == nil {
		return true, nil // ПВЗ с таким ID уже существует
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return false, nil // ПВЗ с таким ID не существует
	}

	// возврат реальной ошибки
	return false, err
}
