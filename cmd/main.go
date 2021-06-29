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
	router.HandleFunc("/", eps.HealthCheck)
	router.HandleFunc("/todos", eps.GetTodos)
	router.HandleFunc("/todos/{todoId}", eps.GetTodo)

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
