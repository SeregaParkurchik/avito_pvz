//go:build integration

package integration

import (
	"avitopvz/config"
	"avitopvz/internal/models"
	"avitopvz/internal/service"
	"avitopvz/internal/storage"
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

type APITestSuite struct {
	suite.Suite

	service service.Interface
	storage storage.Interface

	conn *pgxpool.Pool
}

func (s *APITestSuite) SetupSuite() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	pgConfig, err := config.TestPGConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации PG конфигурации: %v", err)
	}
	fmt.Println(pgConfig.DSN())
	ctx := context.Background()

	conn, err := storage.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}

	s.conn = conn
	s.storage = storage.NewAvitoDB(conn)
	s.service = service.New(s.storage)
}

func (s *APITestSuite) TearDownSuite() {
	ctx := context.Background()

	_, err := s.conn.Exec(ctx, "DELETE FROM products")
	if err != nil {
		log.Printf("Ошибка при удалении данных из таблицы products: %v", err)
	}

	_, err = s.conn.Exec(ctx, "DELETE FROM acceptances")
	if err != nil {
		log.Printf("Ошибка при удалении данных из таблицы acceptances: %v", err)
	}

	_, err = s.conn.Exec(ctx, "DELETE FROM pvz")
	if err != nil {
		log.Printf("Ошибка при удалении данных из таблицы pvz: %v", err)
	}
}

func (s *APITestSuite) TestFullOrderAcceptanceCycle() {
	ctx := context.Background()

	id, err := uuid.FromString("3fa85f64-5717-4562-b3fc-2c963f66afa6")
	if err != nil {
		log.Fatalf("неправильный UUID: %v", err)
	}

	pvz := &models.PVZ{
		ID:               id,
		City:             "Москва",
		RegistrationDate: time.Now(),
	}

	createdPVZ, err := s.service.CreatePVZ(ctx, pvz)
	s.NoError(err)
	s.NotNil(createdPVZ)

	receptions := &models.Receptions{
		PVZID: createdPVZ.ID,
	}
	createdReceptions, err := s.service.CreateReceptions(ctx, receptions)
	s.NoError(err)
	s.NotNil(createdReceptions)

	for i := 0; i < 50; i++ {
		product := &models.Product{
			PVZID: createdReceptions.PVZID,
			Type:  "одежда",
		}
		_, err := s.service.AddProduct(ctx, product)
		s.NoError(err)
	}

	closedReceptions, err := s.service.CloseLastReception(ctx, "3fa85f64-5717-4562-b3fc-2c963f66afa6")
	s.NoError(err)
	s.NotNil(closedReceptions)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
