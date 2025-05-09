package internal

import (
	"net/http"

	"github.com/Xebec19/lms/common/pkg/logger"
	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/db"
	"go.uber.org/zap"
)

func CreateServer() *http.Server {

	conn, err := db.GetDB()
	if err != nil {
		logger.Log.Error("Error getting database connection", zap.Error(err))
		return nil
	}

	if err := db.Migrate(conn); err != nil {
		logger.Log.Error("Error migrating database", zap.Error(err))
		return nil
	}

	r := GetRoutes()

	return &http.Server{
		Addr:    ":" + utils.GetConfig().Port,
		Handler: r,
	}
}
