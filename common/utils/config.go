package utils

import (
	"os"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	Port string `mapstructure:"PORT"`
}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		logger.Log.Error("Error loading .env file", zap.Error(err))
	}

	logger.Log.Info("Loaded .env file")
}

func GetConfig() *Config {

	return &Config{
		Port: os.Getenv("PORT"),
	}
}
