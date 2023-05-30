package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_Map(t *testing.T) {
	deadlineInput := "2023-06-01T10:00:00Z"
	deadline, err := time.Parse(time.RFC3339, deadlineInput)
	if err != nil {
		t.Fatal(err)
	}

	createdAt := time.Now()
	task := Task{
		ID:          "1",
		Name:        "Sample Task",
		CreatedAt:   createdAt,
		Status:      "Pending",
		Description: "Task description",
		Deadline:    deadline,
		Priority:    2,
		Assignee:    "Ali",
		//sdada
	}

	expectedMap := map[string]interface{}{
		"ID":          "1",
		"Name":        "Sample Task",
		"CreatedAt":   createdAt,
		"Status":      "Pending",
		"Description": "Task description",
		"Deadline":    deadline,
		"Priority":    2,
		"Assignee":    "Ali",
	}

	taskMap := task.Map()

	if !reflect.DeepEqual(taskMap, expectedMap) {
		t.Errorf("Map() returned unexpected map.\nExpected: %+v\nGot: %+v", expectedMap, taskMap)
	}
}

func TestTask_Names(t *testing.T) {
	deadlineInput := "2023-06-01T10:00:00Z"
	deadline, err := time.Parse(time.RFC3339, deadlineInput)
	if err != nil {
		t.Fatal(err)
	}

	task := Task{
		ID:          "1",
		Name:        "Sample Task",
		CreatedAt:   time.Now(),
		Status:      "Pending",
		Description: "Task description",
		Deadline:    deadline,
		Priority:    2,
		Assignee:    "Ali",
	}

	expectedNames := []string{
		"ID",
		"Name",
		"CreatedAt",
		"Status",
		"Description",
		"Deadline",
		"Priority",
		"Assignee",
	}

	names := task.Names()

	if !reflect.DeepEqual(names, expectedNames) {
		t.Errorf("Names() returned unexpected field names.\nExpected: %+v\nGot: %+v", expectedNames, names)
	}
}
