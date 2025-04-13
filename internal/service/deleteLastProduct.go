package service

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

func (s *service) DeleteLastProduct(ctx context.Context, pvzIDStr string) error {
	pvzID, err := uuid.FromString(pvzIDStr)
	if err != nil {
		return errors.New("неверный ID ПВЗ")
	}

	err = s.storage.DeleteLastProduct(ctx, pvzID)
	if err != nil {
		return err
	}

	return nil
}
