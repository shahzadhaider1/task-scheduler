package handlers

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"time"

	gen "github.com/shahzadhaider1/task-scheduler/gen/models"
	domain "github.com/shahzadhaider1/task-scheduler/models"
)

// toTaskDomain converts gen to domain model
func toTaskDomain(task *gen.Task) *domain.Task {
	return &domain.Task{
		ID:        task.ID,
		Name:      swag.StringValue(task.Name),
		CreatedAt: time.Time(task.CreatedAt),
		Status:    task.Status,
		Data:      task.Data,
	}
}

// toTaskGen converts domain to gen model
func toTaskGen(task *domain.Task) *gen.Task {
	return &gen.Task{
		ID:        task.ID,
		Name:      swag.String(task.Name),
		CreatedAt: strfmt.DateTime(task.CreatedAt),
		Status:    task.Status,
		Data:      task.Data,
	}
}
