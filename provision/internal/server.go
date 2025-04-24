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

	r.Use(middlewares.LoggingMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	handleRoutes(api)
	handleAuthRoutes(api)

	return &http.Server{
		Addr:    ":" + utils.GetConfig().Port,
		Handler: r,
	}
}

func handleRoutes(api *mux.Router) {
	api.HandleFunc("/health", handler.HandleHealthCheck).Methods("GET")

}

func handleAuthRoutes(api *mux.Router) {
	auth := api.PathPrefix("/auth").Subrouter()

	auth.HandleFunc("/signup", handler.HandleSignup).Methods("POST")
}
