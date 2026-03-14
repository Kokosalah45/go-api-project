package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DBConfig struct {
	Host    string `env:"DB_HOST,required"`
	Port    string `env:"DB_PORT" envDefault:"27017"`
	AppUser string `env:"DB_USERNAME,required"`
	AppPass string `env:"DB_PASSWORD,required"`
	DBName  string `env:"DB_NAME,required"`
}

type AppConfig struct {
	Port   int `env:"PORT" envDefault:"8080"`
	AppEnv string `env:"APP_ENV" envDefault:"local"`
}

type Config struct {
	DB  DBConfig
	App AppConfig
}


func Load() (*Config, error) {
    cfg := &Config{}
    if err := env.Parse(cfg); err != nil {
        return nil, fmt.Errorf("loading config: %w", err)
    }
    return cfg, nil
}