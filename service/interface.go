package service

import (
	"context"

	"github.com/shahzadhaider1/task-scheduler/db"
)

// Service initializes our database instance
type Service struct {
	db db.DataStore
}

// NewService creates a connection to our database
func NewService(ds db.DataStore) *Service {
	return &Service{db: ds}
}

// Close disconnects the mongo client
func (s *Service) Close(ctx context.Context) error {
	return s.db.Disconnect(ctx)
}
