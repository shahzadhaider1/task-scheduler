package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"

	"github.com/shahzadhaider1/task-scheduler/config"
	"github.com/shahzadhaider1/task-scheduler/db"
	domainErr "github.com/shahzadhaider1/task-scheduler/errors"
	"github.com/shahzadhaider1/task-scheduler/models"
)

const (
	taskTableName = "task"
)

func init() {
	db.Register("mysql", NewClient)
}

//The first implementation.
type client struct {
	db *sqlx.DB
}

// formatDSN has database connection variables
func formatDSN() string {
	cfg := mysql.NewConfig()
	cfg.Net = "tcp"
	cfg.Addr = fmt.Sprintf("%s:%s", viper.GetString(config.DbHost), viper.GetString(config.DbPort))
	cfg.DBName = viper.GetString(config.DbName)
	cfg.ParseTime = true
	cfg.User = viper.GetString(config.DbUser)
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
	return &client{db: cli}, nil
}

// AddTask adds task to the database
func (c *client) AddTask(ctx context.Context, task *models.Task) (string, error) {
	if task.ID != "" {
		return "", errors.New("id is not empty")
	}
	task.ID = uuid.NewV4().String()

	names := task.Names()
	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES(%s)`, taskTableName, strings.Join(names, ","), strings.Join(mkPlaceHolder(names, ":", func(name, prefix string) string {
		return prefix + name
	}), ","))
	if _, err := c.db.NamedExec(query, task); err != nil {
		fmt.Println(query)
		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

// GetTask gets the task from database based on id
func (c *client) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var task models.Task
	if err := c.db.Get(&task, fmt.Sprintf(`SELECT * FROM %s WHERE id = '%s'`, taskTableName, id)); err != nil {
		if err == sql.ErrNoRows {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}
		return nil, err
	}

	return &task, nil
}

// UpdateTask updates the task in the database
func (c *client) UpdateTask(ctx context.Context, task *models.Task) error {
	names := task.Names()
	if _, err := c.db.NamedExec(fmt.Sprintf(`UPDATE %s SET %s WHERE id=:id`, taskTableName, strings.Join(mkPlaceHolder(names[1:], "=:", func(name, prefix string) string {
		return name + prefix + name
	}), ",")), task); err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	return nil
}

// DeleteTask deletes the task from database based on id
func (c *client) DeleteTask(ctx context.Context, id string) error {
	if _, err := c.db.Query(fmt.Sprintf(`DELETE FROM %s WHERE id= '%s'`, taskTableName, id)); err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

// Disconnect - closes the db connections
func (c *client) Disconnect(ctx context.Context) error {
	if err := c.db.Close(); err != nil {
		return errors.Wrap(err, "failed to disconnect mysql client")
	}

	return nil
}

func mkPlaceHolder(names []string, prefix string, formatName func(name, prefix string) string) []string {
	ph := make([]string, len(names))
	for i, name := range names {
		ph[i] = formatName(name, prefix)
	}

	return ph
}
