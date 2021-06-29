package entity

import uuid "github.com/satori/go.uuid"

// User entity model
type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	Address      string    `json:"address"`
	PasswordHash string    `json:"-"`
}
