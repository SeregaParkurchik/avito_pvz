package service

import (
	"avitopvz/internal/models"
	"context"
	"errors"

	"github.com/gofrs/uuid"
)

func (s *service) CloseLastReception(ctx context.Context, pvzIDStr string) (*models.Receptions, error) {
	pvzID, err := uuid.FromString(pvzIDStr)
	if err != nil {
		return nil, errors.New("неверный ID ПВЗ")
	}

	receptions, err := s.storage.CloseLastReception(ctx, pvzID)
	if err != nil {
		return nil, err
	}

	return receptions, nil
}
