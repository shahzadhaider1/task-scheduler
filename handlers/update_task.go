package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzadhaider1/task-scheduler"
	domainError "github.com/shahzadhaider1/task-scheduler/errors"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi/operations/service"
)

// NewUpdateTaskHandler handles request for updating Task
func NewUpdateTaskHandler(ctx context.Context, rt *runtime.Runtime) service.UpdateTaskHandler {
	return &updateTask{
		ctx: ctx,
		rt:  rt,
	}
}

type updateTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle handles the update Task request
func (r *updateTask) Handle(params service.UpdateTaskParams) middleware.Responder {
	logrus.Debugf("request:'updateTask' params: %+v", params)

	task := toTaskDomain(params.Task)
	task.ID = params.Task.ID
	if err := r.rt.Service().UpdateTask(r.ctx, task); err != nil {
		switch apiErr := err.(*domainError.APIError); {
		case apiErr.IsError(domainError.NotFound):
			logrus.Errorf("failed to update Task: error[404]: %+v ", err)

			return service.NewUpdateTaskBadRequest()
		default:
			logrus.Errorf("failed to update Task: error[500]: %+v ", err)

			return service.NewUpdateTaskInternalServerError()
		}
	}

	return service.NewUpdateTaskOK()
}
