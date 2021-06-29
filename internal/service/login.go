package service

import (
	"context"
	"encoding/base64"
	"errors"

	"github.com/rendyfebry/go-user/pkg/entity"
	"github.com/rendyfebry/go-user/pkg/rest"
)

func (s *userService) Login(ctx context.Context, req *rest.LoginRequest) (*entity.User, error) {
	user, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	passEnc := base64.StdEncoding.EncodeToString([]byte(req.Password))
	if passEnc != user.PasswordHash {
		return nil, errors.New("Invalid email or password")
	}

	return user, nil
}
