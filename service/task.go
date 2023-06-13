package service

import (
	"context"
	"fmt"

	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/models"
)

// AddTask adds task into database
func (s *Service) AddTask(ctx context.Context, task *models.Task) error {
	existingTask, _ := s.db.GetTaskByID(ctx, task.ID)
	if existingTask != nil {
		return domainErr.NewAPIError(domainErr.Conflict, fmt.Sprintf("task: %s already exists"))
	}

	return s.db.SaveTask(ctx, task)
}

// ListTasks get task record from database
func (s *Service) ListTasks(ctx context.Context) ([]*models.Task, error) {

	tasks, err := s.db.ListTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

// RemoveTask deletes task from database
func (s *Service) RemoveTask(ctx context.Context, id string) error {
	_, err := s.db.GetTaskByID(ctx, id)
	if err != nil {
		return err
	}

	return s.db.RemoveTask(ctx, id)
}

// GetTaskByID gets task from database by id
func (s *Service) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	task, err := s.db.GetTaskByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// UpdateTask updates task record in database
func (s *Service) UpdateTask(ctx context.Context, task *models.Task) error {

	existingTask, err := s.db.GetTaskByID(ctx, task.ID)
	if err != nil {
		return err
	}
	existingTask.Name = task.Name
	existingTask.CreatedAt = task.CreatedAt
	existingTask.Status = task.Status
	existingTask.Description = task.Description
	existingTask.Deadline = task.Deadline
	existingTask.Priority = task.Priority
	existingTask.Assignee = task.Assignee

	return s.db.SaveTask(ctx, existingTask)

}
