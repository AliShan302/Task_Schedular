package db

import (
	"log"

	"github.com/AliShan302/Task_Schedular/models"
)

// DataStore interface represents the common methods for database operations.
type DataStore interface {
	SaveTask(task *models.Task) error
	GetTaskByID(taskID string) (*models.Task, error)
	RemoveTask(taskID string) error
	ListTasks() ([]*models.Task, error)
}

// Option configuration
type Option struct {
	TestMode bool
}

// DataStoreFactory holds configuration for data store
type DataStoreFactory func(conf Option) (DataStore, error)

var datastoreFactories = make(map[string]DataStoreFactory)

// Register saves data store into a data store factory
func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Fatalf("Datastore factory %s does not exist.", name)
		return
	}
	_, ok := datastoreFactories[name]
	if ok {
		log.Fatalf("Datastore factory %s already registered. Ignoring.", name)
		return
	}
	datastoreFactories[name] = factory
}
