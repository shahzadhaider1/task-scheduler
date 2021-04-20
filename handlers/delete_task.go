package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzadhaider1/task-scheduler"
	domainError "github.com/shahzadhaider1/task-scheduler/errors"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi/operations/service"
)

// NewDeleteTaskHandler handles request for removing Task
func NewDeleteTaskHandler(ctx context.Context, rt *runtime.Runtime) service.DeleteTaskHandler {
	return &deleteTask{
		ctx: ctx,
		rt:  rt,
	}
}

type deleteTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle handles the Remove Task request
func (p *deleteTask) Handle(params service.DeleteTaskParams) middleware.Responder {
	logrus.Debugf("request:'deleteTask' params: %+v", params)

	if err := p.rt.Service().DeleteTask(p.ctx, params.ID); err != nil {
		switch apiErr := err.(*domainError.APIError); {
		case apiErr.IsError(domainError.NotFound):
			logrus.Errorf("Bad Request, failed to delete task: error[400]: %+v ", err)

			return service.NewDeleteTaskNotFound()
		default:
			logrus.Errorf("failed to delete task: error[500]: %+v ", err)

			return service.NewDeleteTaskInternalServerError()
		}
	}

	return service.NewDeleteTaskOK()
}
