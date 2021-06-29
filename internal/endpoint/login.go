package endpoint

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rendyfebry/go-user/internal/service"
	"github.com/rendyfebry/go-user/pkg/rest"
)

// MakeLoginEndpoint ...
func MakeLoginEndpoint(s service.UserService) http.HandlerFunc {
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

		var req rest.LoginRequest
		if err := json.Unmarshal(body, &req); err != nil {
			if err != nil {
				res := rest.ErrorResponse{
					Message:  "Invalid request format",
					HTTPCode: http.StatusBadRequest,
				}

				encodeErr(w, res)
				return
			}
		}

		res, err := s.Login(r.Context(), &req)
		if err != nil {
			res := rest.ErrorResponse{
				Message:  err.Error(),
				HTTPCode: http.StatusForbidden,
			}

			encodeErr(w, res)
			return
		}

		encodeResponse(w, http.StatusOK, res)
	}
}
