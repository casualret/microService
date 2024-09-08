package config

import (
	"fmt"
	"github.com/spf13/viper"
)

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
