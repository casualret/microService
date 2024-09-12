package main

import (
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
	"urlshortener/internal/config"
	"urlshortener/internal/handlers"
	"urlshortener/internal/service"
	"urlshortener/internal/storage"
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

	app, err := service.NewApp(logger, db)
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
