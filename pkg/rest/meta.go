package rest

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"
)

// Meta ...
type Meta struct {
	Now       int64  `json:"now"`
	RequestID string `json:"request_id"`
}

// GenerateMeta ...
func GenerateMeta(ctx context.Context) Meta {
	return Meta{
		Now:       time.Now().Unix(),
		RequestID: uuid.NewV4().String(),
	}
}
