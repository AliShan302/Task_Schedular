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

// NewListTask handles a request for listing tasks
func NewListTask(ctx context.Context, rt *runtime.Runtime) operations.ListTasksHandler {
	return &listTask{
		ctx: ctx,
		rt:  rt,
	}
}

type listTask struct {
	ctx context.Context
	rt  *runtime.Runtime
}

// Handle the list tasks request
func (h *listTask) Handle(params operations.ListTasksParams) middleware.Responder {
	tasks, err := h.rt.Service().ListTasks(h.ctx)
	if err != nil {
		switch apiErr := err.(*domainErr.APIError); {
		case apiErr.IsError(domainErr.NotFound):
			return operations.NewListTasksNotFound()
		default:
			return operations.NewListTasksInternalServerError()
		}
	}

	var payload []*models.Task
	for _, task := range tasks {
		payload = append(payload, &models.Task{
			ID:          task.ID,
			Name:        task.Name,
			CreatedAt:   strfmt.DateTime(task.CreatedAt),
			Status:      task.Status,
			Description: task.Description,
			Deadline:    strfmt.DateTime(task.Deadline),
			Priority:    int64(task.Priority),
			Assignee:    task.Assignee,
		})
	}

	return operations.NewListTasksOK().WithPayload(payload)
}
