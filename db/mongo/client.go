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

	"github.com/shahzadhaider1/task-scheduler/config"
	"github.com/shahzadhaider1/task-scheduler/db"
	domainErr "github.com/shahzadhaider1/task-scheduler/errors"
	"github.com/shahzadhaider1/task-scheduler/models"
)

const (
	stuCollection = "Task"
)

func init() {
	db.Register("mongo", NewClient)
}

// client struct for mongodb client
type client struct {
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

	return &client{conn: cli}, nil
}

// AddTask adds task to the database
func (c *client) AddTask(ctx context.Context, task *models.Task) (string, error) {
	if task.ID != "" {
		return "", errors.New("id is not empty")
	}
	task.ID = uuid.NewV4().String()
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.InsertOne(ctx, task); err != nil {
		return "", errors.Wrap(err, "failed to add task")
	}

	return task.ID, nil
}

// GetTask gets the task from database based on id
func (c *client) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var task *models.Task
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if err := collection.FindOne(ctx, bson.M{"_id": id}).Decode(&task); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, domainErr.NewAPIError(domainErr.NotFound, fmt.Sprintf("task: %s not found", id))
		}
		return nil, err
	}
	return task, nil
}

// DeleteTask deletes the task from database based on id
func (c *client) DeleteTask(ctx context.Context, id string) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.DeleteOne(ctx, bson.M{"_id": id}); err != nil {
		return errors.Wrap(err, "failed to delete task")
	}

	return nil
}

// UpdateTask updates the task in the database
func (c *client) UpdateTask(ctx context.Context, task *models.Task) error {
	collection := c.conn.Database(viper.GetString(config.DbName)).Collection(stuCollection)
	if _, err := collection.UpdateOne(ctx, bson.M{"_id": task.ID}, bson.M{"$set": task}); err != nil {
		return errors.Wrap(err, "failed to update task")
	}

	return nil
}

// Disconnect - closes the db connections
func (c *client) Disconnect(ctx context.Context) error {
	if err := c.conn.Disconnect(ctx); err != nil {
		return errors.Wrap(err, "failed to disconnect mongo client")
	}

	return nil
}
