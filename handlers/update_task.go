package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/AliShan302/Task_Schedular"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
	"github.com/AliShan302/Task_Schedular/models"
)

// NewUpdateTask handles request for updating task
func NewUpdateTask(rt *runtime.Runtime) operations.UpdateTaskHandler {
	return &updateTask{rt: rt}
}

type updateTask struct {
	rt *runtime.Runtime
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
	err := d.rt.Service().AddTask(&task)
	if err != nil {
		log().Errorf("failed to add task: %s", err)
		return operations.NewUpdateTaskInternalServerError()
	}

	return operations.NewAddTaskCreated().WithPayload(params.Task)
	err = d.rt.Service().UpdateTask(&task)
	if err != nil {
		log().Errorf("failed to update task: %s", err)
		return operations.NewUpdateTaskInternalServerError()
	}

	return operations.NewUpdateTaskOK().WithPayload(params.Task)
}
