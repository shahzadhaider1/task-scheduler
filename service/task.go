package service

import (
	"context"

	"github.com/shahzadhaider1/task-scheduler/models"
)

// AddTask adds task into database
func (s *Service) AddTask(ctx context.Context, task *models.Task) (string, error) {
	return s.db.AddTask(ctx, task)
}

// GetTask gets task from database
func (s *Service) GetTask(ctx context.Context, id string) (*models.Task, error) {
	return s.db.GetTask(ctx, id)
}

// UpdateTask updates the task in the database
func (s *Service) UpdateTask(ctx context.Context, task *models.Task) error {
	return s.db.UpdateTask(ctx, task)
}

// DeleteTask deletes task from database
func (s *Service) DeleteTask(ctx context.Context, id string) error {
	_, err := s.db.GetTask(ctx, id)
	if err != nil {
		return err
	}

	return s.db.DeleteTask(ctx, id)
}
