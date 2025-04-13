package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

func (db *AvitoDB) CloseLastReception(ctx context.Context, pvzID uuid.UUID) (*models.Receptions, error) {
	receptionID, err := db.checkUnclosedReception(ctx, pvzID)
	if err != nil {
		return nil, err
	}

	if receptionID == uuid.Nil {
		return nil, errors.New("нет открытой приемки для закрытия")
	}

	query := `
                UPDATE acceptances
                SET status = 'close'
                WHERE id = $1
                RETURNING id, date_time, pvz_id, status
        `

	updatedReception := &models.Receptions{}
	err = db.pool.QueryRow(ctx, query, receptionID).Scan(
		&updatedReception.ID,
		&updatedReception.DateTime,
		&updatedReception.PVZID,
		&updatedReception.Status,
	)

	if err != nil {
		return nil, err
	}

	return updatedReception, nil
}
