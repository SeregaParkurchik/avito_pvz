package storage

import (
	"avitopvz/internal/models"
	"context"
	"errors"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5"
)

func (db *AvitoDB) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	receptionID, err := db.checkUnclosedReception(ctx, product.PVZID)
	if err != nil {
		return nil, fmt.Errorf("failed to check unclosed reception: %w", err)
	}

	if receptionID == uuid.Nil {
		return nil, errors.New("no open reception found for this PVZ")
	}

	product.ReceptionsID = receptionID

	const query = `
		INSERT INTO products (id, date_time, type, pvz_id, acceptance_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, date_time, type, pvz_id, acceptance_id
	`

	var createdProduct models.Product
	err = db.pool.QueryRow(ctx, query,
		product.ID,
		product.DateTime,
		product.Type,
		product.PVZID,
		product.ReceptionsID,
	).Scan(
		&createdProduct.ID,
		&createdProduct.DateTime,
		&createdProduct.Type,
		&createdProduct.PVZID,
		&createdProduct.ReceptionsID,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("product insertion failed, no rows returned")
		}
		return nil, fmt.Errorf("failed to insert product: %w", err)
	}

	return &createdProduct, nil
}
