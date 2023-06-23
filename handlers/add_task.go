package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"

	runtime "github.com/AliShan302/Task_Schedular"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
	"github.com/AliShan302/Task_Schedular/models"
)

// NewAddTask handles request for saving task
func NewAddTask(rt *runtime.Runtime) operations.AddTaskHandler {
	return &addTask{rt: rt}
}

type addTask struct {
	rt *runtime.Runtime
}

// Handle the add task request
func (d *addTask) Handle(params operations.AddTaskParams) middleware.Responder {
	task := models.Task{
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
		return operations.NewAddTaskConflict()
	}

	return operations.NewAddTaskCreated().WithPayload(params.Task)
}
