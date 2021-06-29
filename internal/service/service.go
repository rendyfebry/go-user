package service

import (
	"context"

	"github.com/rendyfebry/go-user/pkg/entity"
)

// UserService interface
type UserService interface {
	HealthCheck(ctx context.Context) (map[string]string, error)
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
}

// userService object
type userService struct {
}

// New will return new userService object
func New() UserService {
	return &userService{}
}

func (s *userService) HealthCheck(ctx context.Context) (map[string]string, error) {
	status := map[string]string{
		"status": "ok",
	}

	return status, nil
}

func (s *userService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users := []*entity.User{
		{
			Name:  "John",
			Email: "email@example.com",
		},
	}

	return users, nil
}

func (s *userService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user := &entity.User{
		Name:  "John",
		Email: "email@example.com",
	}

	return user, nil
}
