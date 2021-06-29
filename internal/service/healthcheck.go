package service

import "context"

func (s *userService) HealthCheck(ctx context.Context) (map[string]string, error) {
	status := map[string]string{
		"status": "ok",
	}

	return status, nil
}
