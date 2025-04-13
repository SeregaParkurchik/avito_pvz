//go:generate mockery --filename service_mock.go --name Interface --inpackage --with-expecter
package service

import (
	"avitopvz/internal/models"
	"avitopvz/internal/storage"
	"context"
)

type Interface interface {
	Register(ctx context.Context, newUser *models.User) (string, error)
	Login(ctx context.Context, newUser *models.User) (string, error)
	CreatePVZ(ctx context.Context, newPVZ *models.PVZ) (*models.PVZ, error)
	CreateReceptions(ctx context.Context, newReceptions *models.Receptions) (*models.Receptions, error)
	AddProduct(ctx context.Context, product *models.Product) (*models.Product, error)
	CloseLastReception(ctx context.Context, pvzIDStr string) (*models.Receptions, error)
	DeleteLastProduct(ctx context.Context, pvzIDStr string) error
	GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error)
}

type service struct {
	storage storage.Interface
}

func New(storage storage.Interface) Interface {
	return &service{
		storage: storage,
	}
}
