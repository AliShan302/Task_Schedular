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

// NewAddTask handles request for saving task
func NewAddTask(ctx context.Context, rt *runtime.Runtime) operations.AddTaskHandler {
	return &addTask{
		ctx: ctx,
		rt:  rt,
	}
}

type addTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the add task request
func (d *addTask) Handle(params operations.AddTaskParams) middleware.Responder {
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

	if err := d.rt.Service().AddTask(d.ctx, &task); err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.Conflict):
			return operations.NewAddTaskConflict()
		default:
			return operations.NewAddTaskInternalServerError()
		}
	}

	return operations.NewAddTaskCreated().WithPayload(params.Task)
}
