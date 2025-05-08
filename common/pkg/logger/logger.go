package logger

import (
	"os"

	"go.uber.org/zap"
)

var Log *zap.Logger

func Init() {

	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "development" // Default to development if not set
	}

	isProd := env == "production"

	if isProd {
		// Production logger configuration
		Log, _ = zap.NewProduction()
	} else {
		// Development logger configuration
		Log, _ = zap.NewDevelopment()
	}

}

func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}
