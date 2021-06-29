package endpoint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"
	"github.com/rendyfebry/go-user/pkg/entity"
	"github.com/rendyfebry/go-user/pkg/rest"

	"github.com/gorilla/mux"
)

// MakeCreateUserEndpoint ...
func MakeCreateUserEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
		if err != nil {
			res := rest.ErrorResponse{
				Message:  err.Error(),
				HTTPCode: http.StatusInternalServerError,
			}

			encodeErr(w, res)
			return
		}

		if err := r.Body.Close(); err != nil {
			if err != nil {
				res := rest.ErrorResponse{
					Message:  err.Error(),
					HTTPCode: http.StatusInternalServerError,
				}

				encodeErr(w, res)
				return
			}
		}

		var user entity.User
		if err := json.Unmarshal(body, &user); err != nil {
			if err != nil {
				res := rest.ErrorResponse{
					Message:  "Invalid request format",
					HTTPCode: http.StatusBadRequest,
				}

				encodeErr(w, res)
				return
			}
		}

		res, _ := s.CreateUser(r.Context(), &user)

		encodeResponse(w, http.StatusOK, res)
	}
}

// MakeGetUsersEndpoint ...
func MakeGetUsersEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, _ := s.GetUsers(r.Context())

		res := rest.GetUsersResponse{
			Data: users,
			Meta: rest.GenerateMeta(r.Context()),
		}

		encodeResponse(w, http.StatusOK, res)
	}
}

// MakeGetUserEndpoint ...
func MakeGetUserEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		userID := vars["userID"]

		user, err := s.GetUser(r.Context(), userID)
		if err != nil {
			res := rest.ErrorResponse{
				Message:  err.Error(),
				HTTPCode: http.StatusInternalServerError,
			}

			if err.Error() == "Not found" {
				res.HTTPCode = http.StatusNotFound
			}

			encodeErr(w, res)
			return
		}

		res := rest.GetUserResponse{
			Data: user,
			Meta: rest.GenerateMeta(r.Context()),
		}

		encodeResponse(w, http.StatusOK, res)
	}
}
