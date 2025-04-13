package service

import (
	"avitopvz/internal/models"
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

func (s *service) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	if product.Type != "электроника" && product.Type != "одежда" && product.Type != "обувь" {
		return nil, errors.New("такого типа товаров нет")
	}

	product.ID, _ = uuid.NewV4()
	product.DateTime = time.Now().UTC()

	addProduct, err := s.storage.AddProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	utcDateTime := addProduct.DateTime.UTC()
	addProduct.DateTime = utcDateTime

	return addProduct, nil
}
