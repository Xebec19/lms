package middlewares

import (
	"net/http"

	"github.com/Xebec19/lms/common/pkg/logger"
	"go.uber.org/zap"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		logger.Log.Info("HTTP request", zap.String("method", r.Method), zap.String("url", r.URL.String()), zap.String("remote", r.RemoteAddr))
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
