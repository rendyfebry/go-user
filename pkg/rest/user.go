package rest

import (
	"github.com/rendyfebry/go-user/pkg/entity"
)

type (
	// CreateUserRequest ...
	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		Password string `json:"password"`
	}

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

	// UpdateUserRequest ...
	UpdateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Address  string `json:"address"`
		Password string `json:"password"`
		ID       string `json:"-"`
	}

	// LoginRequest ...
	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
)
