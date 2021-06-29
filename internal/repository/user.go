package repository

import (
	"context"
	"errors"

	"github.com/rendyfebry/go-user/pkg/entity"

	uuid "github.com/satori/go.uuid"
)

func (r *userRepo) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user.ID = uuid.NewV4()
	users = append(users, user)

	return user, nil
}

func (r *userRepo) GetUsers(ctx context.Context) ([]*entity.User, error) {
	return users, nil
}

func (r *userRepo) GetUser(ctx context.Context, id string) (*entity.User, error) {
	for _, v := range users {
		if v.ID.String() == id {
			return v, nil
		}
	}

	return nil, errors.New("Not found")
}
