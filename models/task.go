package models

import (
	"time"

	"github.com/fatih/structs"
)

// Task represents a task in the task scheduler.
type Task struct {
	ID          string    `json:"id" bson:"_id" db:"id" `
	Name        string    `json:"name" bson:"name" db:"name" `
	CreatedAt   time.Time `json:"created_at" bson:"created_at" db:"created_at"  `
	Status      string    `json:"status" bson:"status" db:"status" `
	Description string    `json:"description" bson:"description" db:"description" `
	Deadline    time.Time `json:"deadline" bson:"deadline" db:"deadline" `
	Priority    int       `json:"priority" bson:"priority" db:"priority" `
	Assignee    string    `json:"assignee" bson:"assignee" db:"assignee" `
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
