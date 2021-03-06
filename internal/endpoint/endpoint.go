package endpoint

import (
	"encoding/json"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"
	"github.com/rendyfebry/go-user/pkg/rest"
)

const contentType = "application/json;charset=UTF-8"

// Endpoints ...
type Endpoints struct {
	HealthCheck http.Handler
	CreateUser  http.Handler
	GetUsers    http.Handler
	GetUser     http.Handler
	UpdateUser  http.Handler
	DeleteUser  http.Handler
	Login       http.Handler
}

// New will return new userService object
func New(s service.UserService) Endpoints {
	// Create Endpoint handler
	eps := Endpoints{
		HealthCheck: MakeHealthCheckEndpoint(s),
		CreateUser:  MakeCreateUserEndpoint(s),
		GetUsers:    MakeGetUsersEndpoint(s),
		GetUser:     MakeGetUserEndpoint(s),
		UpdateUser:  MakeUpdateUserEndpoint(s),
		DeleteUser:  MakeDeleteUserEndpoint(s),
		Login:       MakeLoginEndpoint(s),
	}

	// Apply middleware
	eps.CreateUser = Logger(eps.CreateUser)
	eps.GetUsers = Logger(eps.GetUsers)
	eps.GetUser = Logger(eps.GetUser)
	eps.UpdateUser = Logger(eps.UpdateUser)
	eps.DeleteUser = Logger(eps.DeleteUser)
	eps.Login = Logger(eps.Login)

	return eps
}

func encodeResponse(w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)

	if statusCode != http.StatusNoContent {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	}
}

func encodeErr(w http.ResponseWriter, err rest.ErrorResponse) {
	encodeResponse(w, err.HTTPCode, err)
}
