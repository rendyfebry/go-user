package service

import (
	"context"

	"github.com/rendyfebry/go-user/internal/repository"
	"github.com/rendyfebry/go-user/pkg/entity"
	"github.com/rendyfebry/go-user/pkg/rest"
)

// UserService interface
type UserService interface {
	HealthCheck(ctx context.Context) (string, error)
	CreateUser(ctx context.Context, req *rest.CreateUserRequest) (*entity.User, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
	UpdateUser(ctx context.Context, user *rest.UpdateUserRequest) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
	Login(ctx context.Context, req *rest.LoginRequest) (*entity.User, error)
}

// userService object
type userService struct {
	Repo repository.UserRepository
}

// New will return new userService object
func New() UserService {
	return &userService{
		Repo: repository.New(),
	}
}
