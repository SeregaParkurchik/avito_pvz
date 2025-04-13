//go:generate mockery --filename sorage_mock.go --name Interface --inpackage --with-expecter
package storage

import (
	"avitopvz/internal/models"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Interface interface {
	Register(ctx context.Context, newUser *models.User) error
	Login(ctx context.Context, user *models.User) (*models.User, error)
	CreatePVZ(ctx context.Context, newPVZ *models.PVZ) (*models.PVZ, error)
	CreateReceptions(ctx context.Context, newReceptions *models.Receptions) (*models.Receptions, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	CloseLastReception(ctx context.Context, pvzID uuid.UUID) (*models.Receptions, error)
	DeleteLastProduct(ctx context.Context, pvzIDStr uuid.UUID) error
	GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error)
}

type AvitoDB struct {
	pool *pgxpool.Pool
}

func NewAvitoDB(pool *pgxpool.Pool) *AvitoDB {
	return &AvitoDB{pool: pool}
}

func New(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error parsing connection string: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %w", err)
	}

	return pool, nil
}
