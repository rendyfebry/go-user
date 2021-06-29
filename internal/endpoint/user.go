package endpoint

import (
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rendyfebry/go-user/internal/service"
	"github.com/rendyfebry/go-user/pkg/entity"
)

// MakeCreateUserEndpoint ...
func MakeCreateUserEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			panic(err)
		}

		if err := r.Body.Close(); err != nil {
			panic(err)
		}

		var user entity.User
		if err := json.Unmarshal(body, &user); err != nil {
			panic(err)
		}

		res, _ := s.CreateUser(context.Background(), &user)

		encodeResponse(w, res)
	}
}

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
