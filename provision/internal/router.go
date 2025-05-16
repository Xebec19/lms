package internal

import (
	"github.com/Xebec19/lms/common/middlewares"
	"github.com/Xebec19/lms/provision/internal/handler"
	"github.com/gorilla/mux"
)

func GetRoutes() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggingMiddleware)

	api := r.PathPrefix("/api").Subrouter()

	handleRoutes(api)
	handleUserRoutes(api)

	return r
}

func handleRoutes(api *mux.Router) {
	api.HandleFunc("/health", handler.HandleHealthCheck).Methods("GET")

}

func handleUserRoutes(api *mux.Router) {
	auth := api.PathPrefix("/v1/user/auth").Subrouter()

	auth.HandleFunc("/signup", handler.HandleSignup).Methods("POST")
}
