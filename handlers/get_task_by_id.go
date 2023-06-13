package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	runtime "github.com/AliShan302/Task_Schedular"
	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/gen/models"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
)

// NewGetTaskByID handles a request for retrieving Task
func NewGetTaskByID(ctx context.Context, rt *runtime.Runtime) operations.GetTaskByIDHandler {
	return &getTask{
		ctx: ctx,
		rt:  rt,
	}
}

type getTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the get Task request
func (d *getTask) Handle(params operations.GetTaskByIDParams) middleware.Responder {
	task, err := d.rt.Service().GetTaskByID(d.ctx, params.ID)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewGetTaskByIDNotFound()
		default:
			return operations.NewGetTaskByIDInternalServerError()
		}
	}

	return operations.NewGetTaskByIDOK().WithPayload(&models.Task{
		Name:        task.Name,
		CreatedAt:   strfmt.DateTime(task.CreatedAt),
		Status:      task.Status,
		Description: task.Description,
		Deadline:    strfmt.DateTime(task.Deadline),
		Priority:    int64(task.Priority),
		Assignee:    task.Assignee,
	})
}
