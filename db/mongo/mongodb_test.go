package mongo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/AliShan302/Task_Schedular/db"
	"github.com/AliShan302/Task_Schedular/models"
)

func Test_mongoClient_GetTaskByID(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")

	m, _ := NewClient(db.Option{})
	task := &models.Task{

		Name:      "Sample Task",
		CreatedAt: time.Now(),
		Status:    "Pending",
		Deadline:  time.Now(),
		Priority:  1,
		Assignee:  "ali",
	}
	err := m.SaveTask(task)
	if err != nil {
		t.Errorf("Failed to save task: %v", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - get task from db",
			args:    args{task.ID},
			wantErr: false,
		},
		{
			name:    "fail - Invalid id",
			args:    args{id: "3256a0faf-dcadf-4aced-9fa18-fea88f4760724"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			got, err := m.GetTaskByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Printf("got: %#v\n", got)
		})
	}
}

func Test_mongoClient_ListTasks(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	type args struct {
		task *models.Task
	}
	m, _ := NewClient(db.Option{})
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name: "success - List Tasks from db",

			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {

			got, err := m.ListTasks()
			if (err != nil) != testCase.wantErr {
				t.Errorf("ListTasks() error = %v, wantErr %v", err, testCase.wantErr)
			}

			for _, task := range got {
				fmt.Printf("Task ID: %s\n", task.ID)
				fmt.Printf("Task Name: %s\n", task.Name)
				fmt.Printf("Task Description: %s\n", task.Description)
				// Print other task fields as needed
				fmt.Println()
			}
		})
	}
}

func Test_mongoClient_RemoveTask(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	m, _ := NewClient(db.Option{})
	task := &models.Task{

		Name:      "Sample Task",
		CreatedAt: time.Now(),
		Status:    "Pending",
		Deadline:  time.Now(),
		Priority:  1,
		Assignee:  "ali",
	}
	err := m.SaveTask(task)
	if err != nil {
		t.Errorf("Failed to save task: %v", err)
	}

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - Remove task from db",
			args:    args{task.ID},
			wantErr: false,
		},
		{
			name:    "error - non-existent task ID",
			args:    args{id: "non-existent-id"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := m.RemoveTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveTask() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}

func Test_mongoClient_SaveTask(t *testing.T) {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "localhost")
	deadlineInput := "2023-06-01T10:00:00Z"
	deadline, err := time.Parse(time.RFC3339, deadlineInput)
	if err != nil {
		t.Fatal(err)
	}
	createdAt := time.Now()
	type args struct {
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name: "success - add task in db",
			args: args{task: &models.Task{
				Name:        "Sample Task",
				CreatedAt:   createdAt,
				Status:      "Pending",
				Description: "Task description",
				Deadline:    deadline,
				Priority:    2,
				Assignee:    "Ali"}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, _ := NewClient(db.Option{})
			err := m.SaveTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveTask() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
