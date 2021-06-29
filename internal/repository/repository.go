package repository

import (
	"context"

	"github.com/rendyfebry/go-user/pkg/entity"
)

var users []*entity.User

func init() {
	users = []*entity.User{
		{
			Name:  "John",
			Email: "email@example.com",
		},
		{
			Name:  "Doe",
			Email: "doe@example.com",
		},
	}
}

// UserRepository interface
type UserRepository interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUser(ctx context.Context, id string) (*entity.User, error)
}

// userRepo object
type userRepo struct {
	// We can put DB connection here
}

// New will return new userRepo object
func New() UserRepository {
	return &userRepo{}
}

func (r *userRepo) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return users, nil
}

func (r *userRepo) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return users[0], nil
}
