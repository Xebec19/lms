package utils

import (
	"os"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	Port        string `mapstructure:"PORT"`
	DB_HOST     string `mapstructure:"DB_HOST"`
	DB_PORT     string `mapstructure:"DB_PORT"`
	DB_NAME     string `mapstructure:"DB_NAME"`
	DB_USER     string `mapstructure:"DB_USER"`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD"`
	GO_ENV      string `mapstructure:"GO_ENV"`
	JWT_SECRET  string `mapstructure:"JWT_SECRET"`
	APP_NAME    string `mapstructure:"APP_NAME"`
	APP_URL     string `mapstructure:"APP_URL"`
	WEB_URL     string `mapstructure:"WEB_URL"`
}

func InitConfig() {

	if os.Getenv("GO_ENV") == "test" {
		logger.Log.Info("Running in test mode, skipping .env loading")
		return
	}

	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file", zap.Error(err))
	}

	logger.Log.Info("Loaded .env file")
}

func GetConfig() *Config {

	return &Config{
		Port:        os.Getenv("PORT"),
		DB_HOST:     os.Getenv("DB_HOST"),
		DB_PORT:     os.Getenv("DB_PORT"),
		DB_NAME:     os.Getenv("DB_NAME"),
		DB_USER:     os.Getenv("DB_USER"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		GO_ENV:      os.Getenv("GO_ENV"),
		JWT_SECRET:  os.Getenv("JWT_SECRET"),
		APP_NAME:    os.Getenv("APP_NAME"),
		APP_URL:     os.Getenv("APP_URL"),
		WEB_URL:     os.Getenv("WEB_URL"),
	}
}
