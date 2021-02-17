package mongo

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
	client, _ := NewClient(db.Option{})

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
			name: "success - add task in db",
			args: args{task: &models.Task{
				Name:      "AddTask",
				CreatedAt: time.Time{},
				Status:    "active",
				Data:      "Any Data of Save Task",
			}},
			wantErr: false,
		},
		{
			name: "fail - add task in db",
			args: args{task: &models.Task{
				ID:        "1",
				Name:      "AddTask-fail",
				CreatedAt: time.Time{},
				Status:    "active",
				Data:      "Any data of Save Task - Fail",
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.AddTask(tt.args.ctx, tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
		})
	}
}

func Test_client_DeleteTask(t *testing.T) {
	setDBENV()
	client, _ := NewClient(db.Option{})

	task := &models.Task{
		Name:      "DeleteTask",
		CreatedAt: time.Time{},
		Status:    "active",
		Data:      "Any data of Delete Task",
	}
	taskID, _ := client.AddTask(context.TODO(), task)

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
			if err := client.DeleteTask(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetTask(t *testing.T) {
	setDBENV()
	client, _ := NewClient(db.Option{})

	task1 := &models.Task{
		Name:      "GetTask",
		CreatedAt: time.Time{},
		Status:    "active",
		Data:      "Data of Get Task",
	}
	task1ID, _ := client.AddTask(context.TODO(), task1)

	task2 := &models.Task{
		Name:      "GetTask",
		CreatedAt: time.Time{},
		Status:    "active",
		Data:      "Data of Get Task",
	}
	task2ID, _ := client.AddTask(context.TODO(), task2)

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
			name:    "success - get task1 from db",
			args:    args{id: task1ID},
			want:    task1,
			wantErr: false,
		},
		{
			name:    "success - get task2 from db",
			args:    args{id: task2ID},
			want:    task2,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.GetTask(tt.args.ctx, tt.args.id)
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
	client, _ := NewClient(db.Option{})

	task1 := &models.Task{
		Name:      "UpdateTask",
		CreatedAt: time.Time{},
		Status:    "active",
		Data:      "Data of Update Task",
	}
	_, _ = client.AddTask(context.TODO(), task1)

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
			name:    "success - update task in db",
			args:    args{task: task1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := client.UpdateTask(tt.args.ctx, tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// setDBENV has connection for DB
func setDBENV() {
	os.Setenv("DB_PORT", "27017")
	os.Setenv("DB_HOST", "task-scheduler-mongo-db")
}
