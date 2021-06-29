package service

import (
	"context"
	"fmt"
)

// UserService interface
type UserService interface {
	HealthCheck(ctx context.Context) (string, error)
	GetTodos(ctx context.Context) (string, error)
	GetTodo(ctx context.Context, id string) (string, error)
}

// userService object
type userService struct {
}

// New will return new userService object
func New() UserService {
	return &userService{}
}

func (s *userService) HealthCheck(ctx context.Context) (string, error) {
	return "healthcheck", nil
}

func (s *userService) GetTodos(ctx context.Context) (string, error) {
	return "GetTodos", nil
}

func (s *userService) GetTodo(ctx context.Context, id string) (string, error) {
	return fmt.Sprintf("Get Todo: %s", id), nil
}
