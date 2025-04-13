package main

import (
	"avitopvz/config"
	"avitopvz/internal/api"
	"avitopvz/internal/routes"
	"avitopvz/internal/service"
	"avitopvz/internal/storage"
	"avitopvz/pkg/pvz_v1"
	"context"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	httpConfig, err := config.NewHTTPConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации HTTP конфигурации: %v", err)
	}

	pgConfig, err := config.NewPGConfig()
	if err != nil {
		log.Fatalf("Ошибка инициализации PG конфигурации: %v", err)
	}

	ctx := context.Background()
	dbConn, err := storage.New(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbConn.Close()

	dbStorage := storage.NewAvitoDB(dbConn)

	httpService := service.New(dbStorage)

	userHandler := api.NewUserHandler(httpService)

	router := routes.InitRoutes(userHandler)

	address := httpConfig.Address()

	log.Printf("Запуск HTTP сервера на %s", address)
	go func() {
		log.Fatal(router.Run(address))
	}()

	// gRPC сервер
	grpcAddress := ":3000"
	lis, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		log.Fatalf("не удалось прослушать: %v", err)
	}

	grpcServer := grpc.NewServer()
	pvzServer := api.NewPVZServer(dbConn)
	pvz_v1.RegisterPVZServiceServer(grpcServer, pvzServer)
	reflection.Register(grpcServer)

	log.Printf("Запуск gRPC сервера на %s", grpcAddress)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("не удалось запустить gRPC сервер: %v", err)
	}
}
