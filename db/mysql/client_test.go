package mysql

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/shahzadhaider1/task-scheduler/db"
	"github.com/shahzadhaider1/task-scheduler/models"
)

func Test_client_AddTask(t *testing.T) {
	setDBENV()

	type args struct {
		ctx  context.Context
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - add task in db",
			args:    args{task: &models.Task{Name: "AddTask", CreatedAt: time.Now(), Status: "active", Data: "Some data of AddTask"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid task in db",
			args:    args{task: &models.Task{ID: "4", Name: "AddTask-fail", CreatedAt: time.Now(), Status: "active", Data: "Some data of AddTask-fail"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}

func Test_client_DeleteTask(t *testing.T) {
	setDBENV()

	c, _ := NewClient(db.Option{})
	task := &models.Task{Name: "DeleteTask", CreatedAt: time.Now(), Status: "active", Data: "Data of DeleteTask"}
	taskID, _ := c.AddTask(context.TODO(), task)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - delete task from db",
			args:    args{id: taskID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteTask(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetTask(t *testing.T) {
	setDBENV()

	c, _ := NewClient(db.Option{})
	task := &models.Task{Name: "GetTask", CreatedAt: time.Now().UTC().Truncate(time.Minute), Status: "active", Data: "Data of GetTask"}
	taskID, _ := c.AddTask(context.TODO(), task)
	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Task
		wantErr bool
	}{
		// test cases
		{
			name:    "success - get task from db",
			args:    args{id: taskID},
			want:    task,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetTask(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateTask(t *testing.T) {
	setDBENV()

	c, _ := NewClient(db.Option{})
	type args struct {
		ctx  context.Context
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update task in db",
			args:    args{task: &models.Task{Name: "UpdateTask", CreatedAt: time.Now().UTC().Truncate(time.Minute), Status: "active", Data: "Data of UpdateTask"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// setDBENV has connection for DB
func setDBENV() {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "task-scheduler-mysql-db")
	os.Setenv("DB_USER", "root")
}
