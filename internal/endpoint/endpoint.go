package endpoint

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"

	"github.com/gorilla/mux"
)

const contentType = "application/json;charset=UTF-8"

// Endpoints ...
type Endpoints struct {
	HealthCheck http.HandlerFunc
	GetUsers    http.HandlerFunc
	GetUser     http.HandlerFunc
}

// New will return new userService object
func New(s service.UserService) Endpoints {
	eps := Endpoints{
		HealthCheck: func(w http.ResponseWriter, r *http.Request) {
			res, _ := s.HealthCheck(context.Background())

			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(res); err != nil {
				panic(err)
			}
		},
		GetUsers: func(w http.ResponseWriter, r *http.Request) {
			res, _ := s.GetUsers(context.Background())

			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(res); err != nil {
				panic(err)
			}
		},
		GetUser: func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			userID := vars["userID"]

			res, _ := s.GetUser(context.Background(), userID)

			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(http.StatusOK)

			if err := json.NewEncoder(w).Encode(res); err != nil {
				panic(err)
			}
		},
	}

	return eps
}
