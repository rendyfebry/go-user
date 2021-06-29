package service

import (
	"context"
	"encoding/base64"

	"github.com/rendyfebry/go-user/pkg/entity"
	"github.com/rendyfebry/go-user/pkg/rest"
	uuid "github.com/satori/go.uuid"
)

func (s *userService) CreateUser(ctx context.Context, req *rest.CreateUserRequest) (*entity.User, error) {
	passEnc := base64.StdEncoding.EncodeToString([]byte(req.Password))

	user := &entity.User{
		Name:         req.Name,
		Email:        req.Email,
		Address:      req.Address,
		PasswordHash: passEnc,
	}

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

func (s *userService) UpdateUser(ctx context.Context, req *rest.UpdateUserRequest) (*entity.User, error) {
	_, err := s.Repo.GetUser(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	// For testing we just using simple base64 hash
	passEnc := base64.StdEncoding.EncodeToString([]byte(req.Password))
	user := &entity.User{
		ID:           uuid.FromStringOrNil(req.ID),
		Name:         req.Name,
		Email:        req.Email,
		Address:      req.Address,
		PasswordHash: passEnc,
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
