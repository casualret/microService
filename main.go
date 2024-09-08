package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"urlshortener/internal/config"
)

func main() {

	//load config

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	//setup logger

	//connect database && cash
	db, err := MustNewStorage(cfg)
	if err != nil { //df
		panic(err)
	}

	//init service
	_ = db

	//init handlers && routs
	r := InitRoutes()

	//start server
	err = r.Run("localhost:8081")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

type Postgres struct {
	db *sqlx.DB
}

func MustNewStorage(cfg *config.Config) (*Postgres, error) {
	connInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresCfg.PgHost, cfg.PostgresCfg.PgPort, cfg.PostgresCfg.PgUser, cfg.PostgresCfg.PgPassword, cfg.PostgresCfg.PgDatabase, cfg.PostgresCfg.PgSslmode)
	db, err := sqlx.Connect("postgres", connInfo) // connect to postgres
	if err != nil {
		return nil, fmt.Errorf("postgres connect error: %v", err)
	}
	return &Postgres{db: db}, nil
}

func InitRoutes() *gin.Engine {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Пока я мастерил фрегат мир стал бессмыслено богат и полон гнуси")
	})
	//r.POST("/create_url", CreateShortUrl(c))
	//r.GET("/get_url", RedirectUrl(c))
	return r
}
