package endpoint

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendyfebry/go-user/internal/service"
)

// MakeGetUsersEndpoint ...
func MakeGetUsersEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := s.GetUsers(context.Background())

		encodeResponse(w, res)
	}
}

// MakeGetUserEndpoint ...
func MakeGetUserEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["userID"]

		res, _ := s.GetUser(context.Background(), userID)

		encodeResponse(w, res)
	}
}
