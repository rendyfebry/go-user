package rest

import "github.com/rendyfebry/go-user/pkg/entity"

type (
	// GetUsersResponse ...
	GetUsersResponse struct {
		Data []*entity.User `json:"data"`
		Meta Meta           `json:"meta"`
	}

	// GetUserResponse ...
	GetUserResponse struct {
		Data *entity.User `json:"data"`
		Meta Meta         `json:"meta"`
	}

	// LoginRequest ...
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
