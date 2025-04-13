package service

import (
	"avitopvz/internal/models"
	"context"
	"errors"
)

func (s *service) GetAllPVZ(ctx context.Context, listInfo models.GetAllPVZRequest) ([]models.PVZWithReceptions, error) {
	if listInfo.EndDate.Before(listInfo.StartDate) {
		return nil, errors.New("дата окончания не может быть раньше даты начала")
	}

	if listInfo.Page <= 0 {
		return nil, errors.New("номер страницы должен быть больше 0")
	}

	if listInfo.Limit <= 0 {
		return nil, errors.New("лимит должен быть больше 0")
	}

	return s.storage.GetAllPVZ(ctx, listInfo)
}
