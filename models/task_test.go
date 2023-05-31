package models

import (
	"reflect"
	"testing"
	"time"
)

func TestestCaseask_Map(t *testing.T) {

	deadlineInput := "2023-06-01T10:00:00Z"
	deadline, err := time.Parse(time.RFC3339, deadlineInput)
	if err != nil {
		t.Fatal(err)
	}
	createdAt := time.Now()
	type task struct {
		ID          string
		Name        string
		CreatedAt   time.Time
		Status      string
		Description string
		Deadline    time.Time
		Priority    int
		Assignee    string
	}
	tests := []struct {
		name string
		task task
		want map[string]interface{}
	}{
		{
			name: " success - task struct to map",
			task: task{
				ID:          "1",
				Name:        "Sample Task",
				CreatedAt:   createdAt,
				Status:      "Pending",
				Description: "Task description",
				Deadline:    deadline,
				Priority:    2,
				Assignee:    "Ali",
			},
			want: map[string]interface{}{
				"ID":          "1",
				"Name":        "Sample Task",
				"CreatedAt":   createdAt,
				"Status":      "Pending",
				"Description": "Task description",
				"Deadline":    deadline,
				"Priority":    2,
				"Assignee":    "Ali",
			},
		},
		{
			name: " success - task struct to map with fewer fields",
			task: task{
				Name:      "Sample Task",
				CreatedAt: createdAt,
				Status:    "Pending",
				Deadline:  deadline,
				Priority:  3,
			},
			want: map[string]interface{}{
				"ID":          "",
				"Name":        "Sample Task",
				"CreatedAt":   createdAt,
				"Status":      "Pending",
				"Description": "",
				"Deadline":    deadline,
				"Priority":    3,
				"Assignee":    "",
			},
		},
	}
	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Task{
				ID:          testCase.task.ID,
				Name:        testCase.task.Name,
				CreatedAt:   testCase.task.CreatedAt,
				Status:      testCase.task.Status,
				Description: testCase.task.Description,
				Deadline:    testCase.task.Deadline,
				Priority:    testCase.task.Priority,
				Assignee:    testCase.task.Assignee,
			}
			if got := s.Map(); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Map() = %v, want %v", got, testCase.want)
			}
		})
	}

}

func TestStudent_Names(t *testing.T) {
	deadlineInput := "2023-06-01T10:00:00Z"
	deadline, err := time.Parse(time.RFC3339, deadlineInput)
	if err != nil {
		t.Fatal(err)
	}
	createdAt := time.Now()
	type task struct {
		ID          string
		Name        string
		CreatedAt   time.Time
		Status      string
		Description string
		Deadline    time.Time
		Priority    int
		Assignee    string
	}
	tests := []struct {
		name string
		task task
		want []string
	}{
		{
			name: " success - task struct to map",
			task: task{
				ID:          "1",
				Name:        "Sample Task",
				CreatedAt:   createdAt,
				Status:      "Pending",
				Description: "Task description",
				Deadline:    deadline,
				Priority:    2,
				Assignee:    "Ali",
			},

			want: []string{
				"ID",
				"Name",
				"CreatedAt",
				"Status",
				"Description",
				"Deadline",
				"Priority",
				"Assignee",
			},
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			s := &Task{
				ID:          testCase.task.ID,
				Name:        testCase.task.Name,
				CreatedAt:   testCase.task.CreatedAt,
				Status:      testCase.task.Status,
				Description: testCase.task.Description,
				Deadline:    testCase.task.Deadline,
				Priority:    testCase.task.Priority,
				Assignee:    testCase.task.Assignee,
			}
			if got := s.Names(); !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("Names() = %v, want %v", got, testCase.want)
			}
		})
	}
}
