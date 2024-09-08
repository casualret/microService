package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
)

func main() {

	//load config

	cfg, err := InitConfig()
	if err != nil {
		panic(err)
	}

	//setup logger

	//connect database && cash
	db, err := MustNewStorage(cfg)
	if err != nil {
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

type Config struct {
	HTTPServer struct {
		Address     string `mapstructure:"address"`
		Timeout     string `mapstructure:"timeout"`
		IdleTimeout string `mapstructure:"idle_timeout"`
	} `mapstructure:"http_server"`
	PostgresCfg struct {
		PgHost     string `mapstructure:"pg_host"`
		PgPort     string `mapstructure:"pg_port"`
		PgUser     string `mapstructure:"pg_user"`
		PgPassword string `mapstructure:"pg_password"`
		PgDatabase string `mapstructure:"pg_database"`
		PgSslmode  string `mapstructure:"pg_sslmode"`
	} `mapstructure:"postgres"`
}

func InitConfig() (*Config, error) {
	viper.SetConfigName("config") // config name
	viper.SetConfigType("yaml")   //config type
	viper.AddConfigPath("config") // путь к директории конфигурационного файла

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file: %v", err)
	}
	return &cfg, nil
}

type Postgres struct {
	db *sqlx.DB
}

func MustNewStorage(cfg *Config) (*Postgres, error) {
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
