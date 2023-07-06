package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/AliShan302/Task_Schedular/config"
	"github.com/AliShan302/Task_Schedular/db"
	domainErr "github.com/AliShan302/Task_Schedular/errors"
	"github.com/AliShan302/Task_Schedular/models"
)

const (
	taskTableName = "tasks"
)

func init() {
	db.Register("mysql", NewClient)
}

// The first implementation.
type sqlClient struct {
	db *sqlx.DB
}

// Returns the formatted DSN string.
func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = "root"
	cfg.Passwd = viper.GetString(config.DbPass)
	return cfg.FormatDSN()
}

// NewClient initializes a mysql database connection
func NewClient(conf db.Option) (db.DataStore, error) {
	log().Info("initializing mysql connection: " + formatDSN())
	cli, err := sqlx.Connect("mysql", formatDSN())
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to db")
	}

	return &sqlClient{db: cli}, nil
}

// GetTaskByID get tasks by id from mysql db
func (c *sqlClient) GetTaskByID(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	if err := c.db.Get(&task, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, taskTableName, id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}

		return nil, err
	}

	return &task, nil
}

// SaveTask save tasks in mysql db
func (c *sqlClient) SaveTask(ctx context.Context, task *models.Task) error {
	// Generate a new UUID for the task if ID is empty
	if task.ID == "" {

		task.ID = uuid.NewV4().String()
	}

	// Prepare the SQL query for inserting or updating the task
	query := `
		INSERT INTO tasks ( id, name, created_at, status, description, deadline, priority, assignee)
		VALUES (:id, :name, :created_at, :status, :description, :deadline, :priority, :assignee)

		ON DUPLICATE KEY UPDATE
			name = :name,
			created_at = :created_at,
			status = :status,
			description = :description,
			deadline = :deadline,
			priority = :priority,
			assignee = :assignee
	`

	// Execute the query with named parameters
	if _, err := c.db.NamedExec(query, task); err != nil {
		return errors.Wrap(err, "failed to save a task")
	}

	return nil
}

// RemoveTask removes the tasks from mysql db
func (c *sqlClient) RemoveTask(ctx context.Context, id string) error {
	result, err := c.db.Exec(fmt.Sprintf(`DELETE FROM %s WHERE id= ?`, taskTableName), id)
	if err != nil {
		return errors.Wrap(err, "failed to remove task")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

// ListTasks lists the tasks from mysql db
func (c *sqlClient) ListTasks(ctx context.Context) ([]*models.Task, error) {
	// Prepare the SQL query to retrieve tasks
	query := "SELECT id, name, description, deadline, priority, assignee, created_at, status FROM tasks"

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*models.Task{}
	for rows.Next() {
		task := &models.Task{}

		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.Deadline, &task.Priority, &task.Assignee, &task.CreatedAt, &task.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}
	if len(tasks) == 0 {
		return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found"))
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

// Disconnect - closes the db connections
func (c *sqlClient) Disconnect(ctx context.Context) error {
	if err := c.db.Close(); err != nil {
		return errors.Wrap(err, "failed to disconnect mysql client")
	}

	return nil
}
