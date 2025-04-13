package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *AvitoDB) CreateReceptions(ctx context.Context, newReceptions *models.Receptions) (*models.Receptions, error) {
	existsPVZ, err := db.checkPVZID(ctx, newReceptions.PVZID)
	if err != nil {
		return nil, err
	}
	if !existsPVZ {
		return nil, errors.New("ПВЗ с таким ID не существует")
	}

	existingReceptionID, err := db.checkUnclosedReception(ctx, newReceptions.PVZID)
	if err != nil {
		return nil, err
	}
	if existingReceptionID != uuid.Nil {
		return nil, errors.New("существует незакрытая приемка")
	}

	query := `
			INSERT INTO acceptances (id, date_time, pvz_id, status)
			VALUES ($1, $2, $3, $4)
			RETURNING id, date_time, pvz_id, status
	`

	createdReceptions := &models.Receptions{}
	err = db.pool.QueryRow(ctx, query, newReceptions.ID, newReceptions.DateTime, newReceptions.PVZID, newReceptions.Status).Scan(&createdReceptions.ID, &createdReceptions.DateTime, &createdReceptions.PVZID, &createdReceptions.Status)

	if err != nil {
		return nil, err
	}

	return createdReceptions, nil
}

func (db *AvitoDB) checkUnclosedReception(ctx context.Context, pvzID uuid.UUID) (uuid.UUID, error) {
	query := `
			SELECT id
			FROM acceptances
			WHERE pvz_id = $1 AND status = 'in_progress'
			LIMIT 1
	`

	var receptionID uuid.UUID
	err := db.pool.QueryRow(ctx, query, pvzID).Scan(&receptionID)

	if err == nil {
		return receptionID, nil // Незакрытая приемка найдена, возвращаем ID
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return uuid.Nil, nil // Незакрытая приемка не найдена, возвращаем nil
	}

	// Возврат реальной ошибки
	return uuid.Nil, err
}
