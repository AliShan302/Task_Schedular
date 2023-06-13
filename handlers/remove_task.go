package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/AliShan302/Task_Schedular"
	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
)

// NewDeleteTask function will delete the task
func NewDeleteTask(ctx context.Context, rt *runtime.Runtime) operations.DeleteTaskHandler {
	return &deleteTask{
		ctx: ctx,
		rt:  rt,
	}
}

type deleteTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the delete entry request
func (d *deleteTask) Handle(params operations.DeleteTaskParams) middleware.Responder {
	if err := d.rt.Service().RemoveTask(d.ctx, params.ID); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewDeleteTaskNotFound()
		default:
			return operations.NewDeleteTaskInternalServerError()
		}
	}
	return operations.NewDeleteTaskNoContent()
}
