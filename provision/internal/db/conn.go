package db

import (
	"sync"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/Xebec19/lms/common/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect to database")
		}
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
