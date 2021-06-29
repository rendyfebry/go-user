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
	HealthCheck http.HandlerFunc
	CreateUser  http.HandlerFunc
	GetUsers    http.HandlerFunc
	GetUser     http.HandlerFunc
}

// New will return new userService object
func New(s service.UserService) Endpoints {
	eps := Endpoints{
		HealthCheck: MakeHealthCheckEndpoint(s),
		CreateUser:  MakeCreateUserEndpoint(s),
		GetUsers:    MakeGetUsersEndpoint(s),
		GetUser:     MakeGetUserEndpoint(s),
	}

	return eps
}

func encodeResponse(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func encodeErr(w http.ResponseWriter, err rest.ErrorResponse) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(err.HTTPCode)

	if err := json.NewEncoder(w).Encode(err); err != nil {
		panic(err)
	}
}
