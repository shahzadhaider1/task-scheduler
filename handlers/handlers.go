package handlers

import (
	"context"

	"github.com/go-openapi/loads"

	runtime "github.com/shahzadhaider1/task-scheduler"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.TaskSchedulerAPI

// NewHandler overrides swagger api handlers
func NewHandler(ctx context.Context, rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewTaskSchedulerAPI(spec)

	// Service handlers
	handler.ServiceAddTaskHandler = NewAddTaskHandler(ctx, rt)
	handler.ServiceDeleteTaskHandler = NewDeleteTaskHandler(ctx, rt)
	handler.ServiceGetTaskByIDHandler = NewGetTaskByIDHandler(ctx, rt)
	handler.ServiceUpdateTaskHandler = NewUpdateTaskHandler(ctx, rt)

	return handler
}
