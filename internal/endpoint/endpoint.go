package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"

	"github.com/gorilla/mux"
)

// Endpoints ...
type Endpoints struct {
	HealthCheck http.HandlerFunc
	GetTodos    http.HandlerFunc
	GetTodo     http.HandlerFunc
}

// New will return new userService object
func New(s service.UserService) Endpoints {
	eps := Endpoints{
		HealthCheck: func(w http.ResponseWriter, r *http.Request) {
			res, _ := s.HealthCheck(context.Background())

			fmt.Fprint(w, res)
		},
		GetTodos: func(w http.ResponseWriter, r *http.Request) {
			res, _ := s.GetTodos(context.Background())

			fmt.Fprint(w, res)
		},
		GetTodo: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			todoID := vars["todoId"]

			res, _ := s.GetTodo(context.Background(), todoID)

			fmt.Fprint(w, res)
		},
	}

	return eps
}
