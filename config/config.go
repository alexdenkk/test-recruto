package config

import (
	"os"
)

// Config - структура конфигурации приложения
type Config struct {
	Host string `json:"host"`
}

// Load - загрузка конфига из окружающей среды
func Load() *Config {
	return &Config{
		Host: os.Getenv("HOST"),
	}
}
