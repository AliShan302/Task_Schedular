package handlers

import (
	"context"
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/AliShan302/Task_Schedular"
	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
	"github.com/AliShan302/Task_Schedular/models"
)

// NewUpdateTask handles request for updating task
func NewUpdateTask(ctx context.Context, rt *runtime.Runtime) operations.UpdateTaskHandler {
	return &updateTask{
		ctx: ctx,
		rt:  rt,
	}
}

type updateTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the update task request
func (d *updateTask) Handle(params operations.UpdateTaskParams) middleware.Responder {
	task := models.Task{
		ID:          params.Task.ID,
		Name:        params.Task.Name,
		CreatedAt:   time.Time(params.Task.CreatedAt),
		Status:      params.Task.Status,
		Description: params.Task.Description,
		Deadline:    time.Time(params.Task.Deadline),
		Priority:    int(params.Task.Priority),
		Assignee:    params.Task.Assignee,
	}
	if err := d.rt.Service().UpdateTask(d.ctx, &task); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewUpdateTaskNotFound()
		default:
			return operations.NewUpdateTaskInternalServerError()
		}
	}

	return operations.NewUpdateTaskOK().WithPayload(params.Task)

}
