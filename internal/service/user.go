package service

import (
	"context"

	"github.com/rendyfebry/go-user/pkg/entity"
)

func (s *userService) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := s.Repo.CreateUser(ctx, user)

	return user, err
}

func (s *userService) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := s.Repo.GetUsers(ctx)

	return users, err
}

func (s *userService) GetUser(ctx context.Context, id string) (*entity.User, error) {
	user, err := s.Repo.GetUser(ctx, id)

	return user, err
}

func (s *userService) UpdateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	_, err := s.Repo.GetUser(ctx, user.ID.String())
	if err != nil {
		return nil, err
	}

	user, err = s.Repo.UpdateUser(ctx, user)
	return user, err
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	_, err := s.Repo.GetUser(ctx, id)
	if err != nil {
		return err
	}

	err = s.Repo.DeleteUser(ctx, id)

	return err
}
