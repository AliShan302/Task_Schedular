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

// The first implementation.
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
func (m *mongoClient) SaveTask(ctx context.Context, task *models.Task) error {
	if task.ID == "" {

		task.ID = uuid.NewV4().String()
	}

	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)

	filter := bson.M{"_id": task.ID}
	update := bson.M{"$set": task}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(ctx, filter, update, opts); err != nil {
		if err == mongo.ErrNoDocuments {
			return domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found"))
		}

		return err
	}

	return nil
}

// GetTaskByID get tasks by id from mongo db
func (m *mongoClient) GetTaskByID(ctx context.Context, taskID string) (*models.Task, error) {
	var task *models.Task
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)

	if err := collection.FindOne(ctx, bson.M{"_id": taskID}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", taskID))
		}

		return nil, err
	}

	return task, nil
}

// RemoveTask removes the tasks from mongo db
func (m *mongoClient) RemoveTask(ctx context.Context, taskID string) error {
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	result, err := collection.DeleteOne(ctx, bson.M{"_id": taskID})
	if err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("task with ID %s not found", taskID)
	}

	return nil
}

// ListTasks removes the tasks from mongo db
func (m *mongoClient) ListTasks(ctx context.Context) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	collection := m.conn.Database(viper.GetString(config.DbName)).Collection(taskCollection)
	filter := bson.M{}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve list of tasks")
	}
	defer cursor.Close(ctx)

	for cursor.Next(context.Background()) {
		task := &models.Task{}

		if err := cursor.Decode(task); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found"))
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Disconnect - closes the db connections
func (c *mongoClient) Disconnect(ctx context.Context) error {
	if err := c.conn.Disconnect(ctx); err != nil {
		return errors.Wrap(err, "failed to disconnect mongo client")
	}

	return nil
}
