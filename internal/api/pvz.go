package api

import (
	"avitopvz/internal/models"
	"avitopvz/pkg/pvz_v1"
	"context"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type PVZServer struct {
	pool *pgxpool.Pool
	pvz_v1.UnimplementedPVZServiceServer
}

func NewPVZServer(pool *pgxpool.Pool) *PVZServer {
	return &PVZServer{
		pool: pool,
	}
}

func (s *PVZServer) GetPVZList(ctx context.Context, req *pvz_v1.GetPVZListRequest) (*pvz_v1.GetPVZListResponse, error) {
	query := `SELECT id, registration_date, city FROM pvz`
	rows, err := s.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
	}
	defer rows.Close()

	var pvzs []*models.PVZ
	for rows.Next() {
		var pvz models.PVZ
		var idStr string
		err := rows.Scan(&idStr, &pvz.RegistrationDate, &pvz.City)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}
		pvz.ID, err = uuid.FromString(idStr)
		if err != nil {
			return nil, fmt.Errorf("ошибка парсинга UUID: %w", err)
		}
		pvzs = append(pvzs, &pvz)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации по строкам: %w", err)
	}

	var protoPVZs []*pvz_v1.PVZ
	for _, pvz := range pvzs {
		protoPVZs = append(protoPVZs, &pvz_v1.PVZ{
			Id:               pvz.ID.String(),
			RegistrationDate: timestamppb.New(pvz.RegistrationDate),
			City:             pvz.City,
		})
	}

	return &pvz_v1.GetPVZListResponse{
		Pvzs: protoPVZs,
	}, nil
}
