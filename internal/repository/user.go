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

func (r *userRepo) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	for _, v := range users {
		if v.ID.String() == user.ID.String() {
			v.Name = user.Name
			v.Email = user.Email
			return v, nil
		}
	}

	return nil, errors.New("Not found")
}

func (r *userRepo) DeleteUser(ctx context.Context, id string) error {
	for k, v := range users {
		if v.ID.String() == id {

			// Remove the element at index k from users.
			users[k] = users[len(users)-1] // Copy last element to index i.
			users[len(users)-1] = nil      // Erase last element (write zero value).
			users = users[:len(users)-1]   // Truncate slice.

			break
		}
	}

	return nil
}
