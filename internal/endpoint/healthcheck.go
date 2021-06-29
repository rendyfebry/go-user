package endpoint

import (
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"
	"github.com/rendyfebry/go-user/pkg/rest"
)

// MakeHealthCheckEndpoint ...
func MakeHealthCheckEndpoint(s service.UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status, _ := s.HealthCheck(r.Context())

		res := rest.HealhcheckResponse{
			Data: status,
			Meta: rest.GenerateMeta(r.Context()),
		}

		encodeResponse(w, http.StatusOK, res)
	}
}
