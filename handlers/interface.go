package handlers

import (
	"context"

	"github.com/go-openapi/loads"

	runtime "github.com/AliShan302/Task_Schedular"
	"github.com/AliShan302/Task_Schedular/gen/restapi/operations"
)

// Handler replaces swagger handler
type Handler *operations.TaskSchedularAPI

// NewHandler overrides swagger api handlers
func NewHandler(ctx context.Context, rt *runtime.Runtime, spec *loads.Document) Handler {
	handler := operations.NewTaskSchedularAPI(spec)

	// handlers
	handler.AddTaskHandler = NewAddTask(ctx, rt)
	handler.ListTasksHandler = NewListTask(ctx, rt)
	handler.DeleteTaskHandler = NewDeleteTask(ctx, rt)
	handler.UpdateTaskHandler = NewUpdateTask(ctx, rt)
	handler.GetTaskByIDHandler = NewGetTaskByID(ctx, rt)
	return handler
}
