package db

import (
	"log"
	"os"
	"sync"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/Xebec19/lms/common/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

var (
	db   *gorm.DB
	once sync.Once
)

// GetDB returns a singleton instance of the database connection
func GetDB() (*gorm.DB, error) {
	once.Do(func() {
		var err error
		config := utils.GetConfig()
		dsn := "host=" + config.DB_HOST + " user=" + config.DB_USER + " password=" + config.DB_PASSWORD + " dbname=" + config.DB_NAME + " port=" + config.DB_PORT + " sslmode=disable"

		if config.GO_ENV != "prod" {
			newLogger := gormLogger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags),
				gormLogger.Config{
					LogLevel: gormLogger.Info,
					Colorful: true,
				},
			)

			db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
				Logger: newLogger,
			})

		} else {
			db, err = gorm.Open(postgres.Open(dsn))
		}

		if err != nil {
			panic("failed to connect to database")
		}
		logger.Log.Info("Database connection established")
	})
	return db, nil
}

// CloseDB closes the database connection
func CloseDB() {
	if db != nil {
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to get database instance")
		}
		sqlDB.Close()
		logger.Log.Info("Database connection closed")
	}
}
