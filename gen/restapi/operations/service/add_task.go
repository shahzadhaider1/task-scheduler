// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// AddTaskHandlerFunc turns a function with the right signature into a add task handler
type AddTaskHandlerFunc func(AddTaskParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AddTaskHandlerFunc) Handle(params AddTaskParams) middleware.Responder {
	return fn(params)
}

// AddTaskHandler interface for that can handle valid add task params
type AddTaskHandler interface {
	Handle(AddTaskParams) middleware.Responder
}

// NewAddTask creates a new http.Handler for the add task operation
func NewAddTask(ctx *middleware.Context, handler AddTaskHandler) *AddTask {
	return &AddTask{Context: ctx, Handler: handler}
}

/* AddTask swagger:route POST /internal/tasks service addTask

create task

*/
type AddTask struct {
	Context *middleware.Context
	Handler AddTaskHandler
}

func (o *AddTask) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAddTaskParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
