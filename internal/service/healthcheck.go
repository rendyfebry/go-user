package service

import "context"

func (s *userService) HealthCheck(ctx context.Context) (string, error) {
	return "ok", nil
}
