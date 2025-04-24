package middlewares

import (
	"net/http"

	"github.com/Xebec19/lms/common/pkg/logger"
	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Log.Info("Request", zap.String("method", r.Method),
			zap.String("path", r.URL.Path))

		next.ServeHTTP(w, r)
	})
}
