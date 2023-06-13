package mongo

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/AliShan302/Task_Schedular/config"
	"github.com/AliShan302/Task_Schedular/db"
	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/models"
)

const (
	taskCollection = "tasks"
)

func init() {
	db.Register("mongo", NewClient)
}

//The first implementation.
type mongoClient struct {
	conn *mongo.Client
}

// NewClient initializes a mongo database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	uri := fmt.Sprintf("mongodb://%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	log().Infof("initializing mongodb: %s", uri)
	cli, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &mongoClient{conn: cli}, nil
}

// SaveTask save tasks in mongo db
func (m *mongoClient) SaveTask(task *models.Task) error {
	if task.ID == "" {

		task.ID = uuid.NewV4().String()
	}

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)

	filter := bson.M{"_id": task.ID}
	update := bson.M{"$set": task}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(context.TODO(), filter, update, opts); err != nil {
	if _, err := collection.InsertOne(context.TODO(), task); err != nil {
		return errors.Wrap(err, "failed to save a task")
	}

	return nil

}

// GetTaskByID get tasks by id from mongo db
func (m *mongoClient) GetTaskByID(taskID string) (*models.Task, error) {
	var task *models.Task
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)

	if err := collection.FindOne(context.TODO(), bson.M{"_id": taskID}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", taskID))
		}

		return nil, err
	}
	return task, nil
}

// RemoveTask removes the tasks from mongo db
func (m *mongoClient) RemoveTask(taskID string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	result, err := collection.DeleteOne(context.TODO(), bson.M{"_id": taskID})
	if err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("task with ID %s not found", taskID)
	}

	return nil
}

// ListTasks removes the tasks from mongo db
func (m *mongoClient) ListTasks() ([]*models.Task, error) {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	// Create an empty slice to store the retrieved tasks
	tasks := []*models.Task{}

	// Create an empty filter since we want to retrieve all tasks
	filter := bson.M{}
	// Execute the find operation on the collection
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Iterate over the cursor and decode each document into a Task struct
	for cursor.Next(context.Background()) {
		task := &models.Task{}

		// Decode the current document into the Task struct
		if err := cursor.Decode(task); err != nil {
			return nil, err
		}

		// Append the task to the tasks slice
		tasks = append(tasks, task)
	}

	// Check for any errors during iteration
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
