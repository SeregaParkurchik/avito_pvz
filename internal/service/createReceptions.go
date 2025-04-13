package service

import (
	"avitopvz/internal/models"
	"context"
	"time"

	"github.com/gofrs/uuid"
)

func (s *service) CreateReceptions(ctx context.Context, newReceptions *models.Receptions) (*models.Receptions, error) {
	newReceptions.ID, _ = uuid.NewV4()
	newReceptions.Status = "in_progress"
	newReceptions.DateTime = time.Now().UTC()

	createdReceptions, err := s.storage.CreateReceptions(ctx, newReceptions)
	if err != nil {
		return nil, err
	}

	utcDateTime := createdReceptions.DateTime.UTC()
	createdReceptions.DateTime = utcDateTime

	return createdReceptions, nil
}
