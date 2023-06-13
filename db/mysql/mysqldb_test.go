package mysql

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/AliShan302/Task_Schedular/db"
	"github.com/AliShan302/Task_Schedular/models"
)

func Test_client_SaveTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "task-schedular-mysql-db")
	os.Setenv("DB_USER", "root")
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
			name: "success - save task in db",
			args: args{task: &models.Task{
				Name:        "Sample Task",
				CreatedAt:   createdAt,
				Status:      "Pending",
				Description: "Task description",
				Deadline:    deadline,
				Priority:    2,
				Assignee:    "Ali",
			}},

			wantErr: false,
		},
		{
			name: "fail - add invalid task in db",
			args: args{task: &models.Task{
				ID:          "4",
				Description: "description",
				Deadline:    deadline,
				Priority:    7,
				Assignee:    "shan",
			}},
			wantErr: true,
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			c, err := NewClient(db.Option{})
			if err != nil {
				t.Fatalf("Failed to create task: %v", err)
			}

			err = c.SaveTask(testCase.args.task)
			if (err != nil) != testCase.wantErr {
				t.Errorf("SaveTask() error = %v, wantErr %v", err, testCase.wantErr)
			}

		})
	}
}
func Test_client_RemoveTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "task-schedular-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})

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
			args:    args{id: "62a264ed-c0d6-4cfa-a36d-9a4cc36252e5"},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			if err := c.RemoveTask(testCase.args.id); (err != nil) != testCase.wantErr {
				t.Errorf("RemoveTask() error = %v, wantErr %v", err, testCase.wantErr)
			}

		})
	}
}

func Test_client_GetTaskByID(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "task-schedular-mysql-db")
	os.Setenv("DB_USER", "root")

	c, _ := NewClient(db.Option{})

	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success - get task from db",
			args: args{id: "b9e09f26-9a60-43e7-b2e7-959139c1e47c"},

			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			got, err := c.GetTaskByID(testCase.args.id)
			if (err != nil) != testCase.wantErr {
				t.Errorf("GetTaskByID() error = %v, wantErr %v", err, testCase.wantErr)
				return
			}

			fmt.Printf("got: %#v\n", got)

		})
	}
}

func Test_client_ListTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "task-schedular-mysql-db")
	os.Setenv("DB_USER", "root")

	type args struct {
		task *models.Task
	}
	c, _ := NewClient(db.Option{})

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

			got, err := c.ListTasks()
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
