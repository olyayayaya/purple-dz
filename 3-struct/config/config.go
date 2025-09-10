package config

import (
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

type Config struct {
	Key string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
	color.Red("не удалось найти env файл")
}
	key := os.Getenv("KEY")
	if key == "" {
		panic("не передан параметр кеу в переменные окружения")
	}
	return &Config{
		Key: key,
	}
} 