package service

import (
	"github.com/AliShan302/Task_Schedular/models"
)

// AddTask adds task into database
func (s *Service) AddTask(task *models.Task) error {
	return s.db.SaveTask(task)
}

// ListTasks updates task record in database
func (s *Service) ListTasks() ([]*models.Task, error) {

	tasks, err := s.db.ListTasks()
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// RemoveTask deletes task from database
func (s *Service) RemoveTask(id string) error {
	_, err := s.db.GetTaskByID(id)
	if err != nil {
		return err
	}

	return s.db.RemoveTask(id)
}

// GetTaskByID gets student from database
func (s *Service) GetTaskByID(id string) (*models.Task, error) {
	student, err := s.db.GetTaskByID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

// UpdateTask updates task record in database
func (s *Service) UpdateTask(task *models.Task) error {
	existingTask, err := s.db.GetTaskByID(task.ID)
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

	return s.db.SaveTask(existingTask)
}
