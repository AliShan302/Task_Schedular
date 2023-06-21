package runtime

import (
	"fmt"
	"reflect"
	"testing"

	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/models"
	"github.com/AliShan302/Task_Schedular/service"
)

type mockDataStore struct {
	tasks []*models.Task
}

func (m *mockDataStore) SaveTask(task *models.Task) error {
	m.tasks = append(m.tasks, task)
	return nil
}

func (m *mockDataStore) GetTaskByID(taskID string) (*models.Task, error) {
	for _, task := range m.tasks {
		if task.ID == taskID {
			return task, nil
		}
	}
	return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", taskID))

}

func (m *mockDataStore) RemoveTask(taskID string) error {
	for i, task := range m.tasks {
		if task.ID == taskID {
			m.tasks = append(m.tasks[:i], m.tasks[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID %s not found", taskID)

}

func (m *mockDataStore) ListTasks() ([]*models.Task, error) {

	return m.tasks, nil
}

func TestNewRuntime(t *testing.T) {
	mockSvc := &service.Service{
		Store: &mockDataStore{},
	}

	tests := []struct {
		name    string
		want    *Runtime
		wantErr bool
	}{
		{
			name:    "ValidRuntime",
			want:    &Runtime{svc: mockSvc},
			wantErr: false,
		},
		{
			name: "InvalidRuntime",

			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRuntime()
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRuntime() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRuntime() got = %v, want %v", got, tt.want)
			}
		})
	}
}
