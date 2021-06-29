package main

import (
	"log"
	"net/http"

	"github.com/rendyfebry/go-user/internal/endpoint"
	"github.com/rendyfebry/go-user/internal/service"

	"github.com/gorilla/mux"
)

func main() {
	svc := service.New()
	eps := endpoint.New(svc)

	router := mux.NewRouter().StrictSlash(true)
	router.Methods("GET").Path("/healthcheck").Handler(eps.HealthCheck)
	router.Methods("POST").Path("/users").Handler(eps.CreateUser)
	router.Methods("GET").Path("/users").Handler(eps.GetUsers)
	router.Methods("GET").Path("/users/{userID}").Handler(eps.GetUser)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
