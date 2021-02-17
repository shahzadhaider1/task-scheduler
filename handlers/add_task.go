package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzadhaider1/task-scheduler"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi/operations/service"
)

// NewAddTaskHandler handles request for saving Task
func NewAddTaskHandler(ctx context.Context, rt *runtime.Runtime) service.AddTaskHandler {
	return &addTask{
		ctx: ctx,
		rt:  rt,
	}
}

type addTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle handles the add Task request
func (p *addTask) Handle(params service.AddTaskParams) middleware.Responder {
	logrus.Debugf("request:'addTask' params: %+v", params)

	if _, err := p.rt.Service().AddTask(p.ctx, toTaskDomain(params.Task)); err != nil {
		logrus.Errorf("failed to add task: error[500]: %+v ", err)

		return service.NewAddTaskInternalServerError()
	}

	return service.NewAddTaskCreated().WithPayload(params.Task)
}
