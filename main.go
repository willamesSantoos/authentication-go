package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/willamesSantoos/authentication/internal/configs"
	"github.com/willamesSantoos/authentication/internal/routes"
)

func main() {
	configs.LoadConfig()

	router := mux.NewRouter()
	routes.Initializer(router)

	log.Println(fmt.Sprintf("Server running on port %s", configs.GetServerPort()))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router))
}
