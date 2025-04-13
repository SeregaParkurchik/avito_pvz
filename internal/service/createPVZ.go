package service

import (
	"avitopvz/internal/models"
	"context"
	"errors"
)

func (s *service) CreatePVZ(ctx context.Context, newPVZ *models.PVZ) (*models.PVZ, error) {
	if newPVZ.City != "Москва" && newPVZ.City != "Санкт-Петербург" && newPVZ.City != "Казань" {
		return nil, errors.New("в данном городе нет пвз, либо город указан неверно")
	}

	createdPVZ, err := s.storage.CreatePVZ(ctx, newPVZ)
	if err != nil {
		return nil, err
	}

	utcRegistrationDate := newPVZ.RegistrationDate.UTC()
	createdPVZ.RegistrationDate = utcRegistrationDate

	return createdPVZ, nil
}
