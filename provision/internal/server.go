package internal

import (
	"net/http"

	"github.com/Xebec19/lms/common/middlewares"
	"github.com/Xebec19/lms/common/utils"
	"github.com/Xebec19/lms/provision/internal/handler"
	"github.com/gorilla/mux"
)

func CreateServer() *http.Server {
	r := mux.NewRouter()

	handleRoutes(r)

	r.Use(middlewares.LoggingMiddleware)

	return &http.Server{
		Addr:    ":" + utils.GetConfig().Port,
		Handler: r,
	}
}

func handleRoutes(r *mux.Router) {
	api := r.PathPrefix("/api/v1/provision").Subrouter()

	api.HandleFunc("/health", handler.HandleHealthCheck).Methods("GET")

}

func handleAuthRoutes(r *mux.Router) {
	// auth := r.PathPrefix("/auth").Subrouter()

}
