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
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
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
