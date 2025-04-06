package internal

import (
	"net/http"

	"github.com/Xebec19/lms/users-svc/internal/handler"
	"github.com/Xebec19/lms/users-svc/internal/middlewares"
	"github.com/Xebec19/lms/users-svc/internal/utils"
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
	api := r.PathPrefix("/api/v1/users").Subrouter()

	api.HandleFunc("/health", handler.HandleHealthCheck).Methods("GET")

}
