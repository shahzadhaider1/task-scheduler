package runtime

import (
	"github.com/shahzadhaider1/task-scheduler/db"
	"github.com/shahzadhaider1/task-scheduler/db/mongo"
	"github.com/shahzadhaider1/task-scheduler/service"
)

// Runtime initializes values for entry point to our application
type Runtime struct {
	svc *service.Service
}

// NewRuntime creates a new runtime
func NewRuntime() (*Runtime, error) {
	pService, err := mongo.NewClient(db.Option{})
	if err != nil {
		return nil, err
	}

	return &Runtime{svc: service.NewService(pService)}, err
}

// Service returns pointer to our service variable
func (r Runtime) Service() *service.Service {
	return r.svc
}
