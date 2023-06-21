package handlers

import (
	"github.com/go-openapi/loads"

	runtime "github.com/AliShan302/Task_Schedular"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.TaskSchedularAPI

// NewHandler overrides swagger api handlers
func NewHandler(rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewTaskSchedularAPI(spec)

	// handlers
	handler.AddTaskHandler = NewAddTask(rt)
	handler.ListTasksHandler = NewListTask(rt)
	handler.DeleteTaskHandler = NewDeleteTask(rt)
	handler.UpdateTaskHandler = NewUpdateTask(rt)
	return handler
}
