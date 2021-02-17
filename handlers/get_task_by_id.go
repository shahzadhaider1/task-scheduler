package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/sirupsen/logrus"

	runtime "github.com/shahzadhaider1/task-scheduler"
	domainError "github.com/shahzadhaider1/task-scheduler/errors"
	"github.com/shahzadhaider1/task-scheduler/gen/restapi/operations/service"
)

// NewGetTaskByIDHandler handles request for getting Package by id
func NewGetTaskByIDHandler(ctx context.Context, rt *runtime.Runtime) service.GetTaskByIDHandler {
	return &getTaskByID{
		ctx: ctx,
		rt:  rt,
	}
}

type getTaskByID struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the get Task by ID request
func (p *getTaskByID) Handle(params service.GetTaskByIDParams) middleware.Responder {
	logrus.Debugf("request:'getTaskByID' params: %+v", params)

	task, err := p.rt.Service().GetTask(p.ctx, params.ID)
	if err != nil {
		switch apiErr := err.(*domainError.APIError); {
		case apiErr.IsError(domainError.NotFound):
			logrus.Errorf("failed to get task by ID: error[404]: %+v ", err)

			return service.NewGetTaskByIDNotFound()
		default:
			logrus.Errorf("failed to get task by ID: error[500]: %+v ", err)

			return service.NewGetTaskByIDInternalServerError()
		}
	}

	return service.NewGetTaskByIDOK().WithPayload(toTaskGen(task))
}
