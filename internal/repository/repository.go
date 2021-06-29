package repository

import (
	"context"

	"github.com/rendyfebry/go-user/pkg/entity"

	uuid "github.com/satori/go.uuid"
)

var users []*entity.User

func init() {
	users = []*entity.User{
		{
			ID:    uuid.FromStringOrNil("0674f8ec-027a-4ca1-8388-4c503a2d992d"),
			Name:  "John",
			Email: "email@example.com",
		},
		{
			ID:    uuid.FromStringOrNil("d0465c53-94a4-47c1-8e0f-45eafaa96583"),
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
	UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id string) error
}

// userRepo object
type userRepo struct {
	// We can put DB connection here
}

// New will return new userRepo object
func New() UserRepository {
	return &userRepo{}
}
