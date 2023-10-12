package routes

import (
	"github.com/gorilla/mux"
	"github.com/willamesSantoos/authentication/internal/handlers"
)

func Initializer(router *mux.Router) {
	router.HandleFunc("/api/authenticate", handlers.Authenticate).Methods("POST")

	router.HandleFunc("/api/v1/users", handlers.AuthMiddleware(handlers.GetAll)).Methods("GET")
	router.HandleFunc("/api/v1/users", handlers.AuthMiddleware(handlers.Create)).Methods("POST")
}
