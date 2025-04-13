package storage

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *AvitoDB) DeleteLastProduct(ctx context.Context, pvzID uuid.UUID) error {
	receptionID, err := db.checkUnclosedReception(ctx, pvzID)
	if err != nil {
		return err
	}

	if receptionID == uuid.Nil {
		return errors.New("нет открытой приемки для удаления товара")
	}

	productID, err := db.getLastProductID(ctx, receptionID)
	if err != nil {
		return err
	}

	if productID == uuid.Nil {
		return errors.New("нет товаров для удаления в приемке")
	}

	query := `
                DELETE FROM products
                WHERE id = $1
        `

	_, err = db.pool.Exec(ctx, query, productID)
	if err != nil {
		return err
	}

	return nil
}

func (db *AvitoDB) getLastProductID(ctx context.Context, receptionID uuid.UUID) (uuid.UUID, error) {
	query := `
                SELECT id
                FROM products
                WHERE acceptance_id = $1
                ORDER BY date_time DESC
                LIMIT 1
        `

	var productID uuid.UUID
	err := db.pool.QueryRow(ctx, query, receptionID).Scan(&productID)

	if err == nil {
		return productID, nil
	}

	if errors.Is(err, pgx.ErrNoRows) {
		return uuid.Nil, nil
	}

	return uuid.Nil, err
}
