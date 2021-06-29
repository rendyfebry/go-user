package endpoint

import (
	"context"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"
)

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _ := s.HealthCheck(context.Background())

		encodeResponse(w, res)
	}
}
