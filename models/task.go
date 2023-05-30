package models

import (
	"time"

	"github.com/fatih/structs"
)

// Task represents a task in the task scheduler.
type Task struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"created_at"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Priority    int       `json:"priority"`
	Assignee    string    `json:"assignee"`
}

// Map returns a map representation of the Task struct.
func (t *Task) Map() map[string]interface{} {
	return structs.Map(t)
}

// Names returns the field names of the Task struct.
func (t *Task) Names() []string {
	fields := structs.Fields(t)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}

	return names
}
