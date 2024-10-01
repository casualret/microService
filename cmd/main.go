package main

import (
	"context"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"microService/internal/config"
	"microService/internal/handlers"
	"microService/internal/service"
	"microService/internal/storage"
	"os"
)

func main() {

	//setup logger

	var logger *slog.Logger
	logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	//load config

	cfg, err := config.InitConfig()
	if err != nil {
		logger.Error("error reading config", err)
		panic(err)
	}

	//connect database && cash

	db, err := storage.MustNewStorage(cfg)
	if err != nil {
		logger.Error("error setup storage", err)
		panic(err)
	}

	//init service

	ctx := context.Background()

	app, err := service.NewApp(logger, db, &ctx)
	if err != nil {
		logger.Error("error setup application", err)
		panic(err)
	}

	//init handlers && routs

	handls := handlers.NewHandlers(app, logger)
	router := handls.InitRoutes()

	//start server

	err = router.Run("localhost:8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
