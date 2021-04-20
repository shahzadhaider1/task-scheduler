package models

import (
	"reflect"
	"testing"
	"time"
)

func TestTask_Map(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		CreatedAt time.Time
		Status    string
		Data      string
	}
	var tests = []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: "success - convert task struct to map",
			fields: fields{
				ID:        "1",
				Name:      "Shahzad",
				CreatedAt: time.Now().UTC().Truncate(time.Minute),
				Status:    "active",
				Data:      "convert task struct to map",
			},
			want: map[string]interface{}{
				"id":         "1",
				"name":       "Shahzad",
				"created_at": time.Now().UTC().Truncate(time.Minute),
				"status":     "active",
				"data":       "convert task struct to map",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Task{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				Status:    tt.fields.Status,
				Data:      tt.fields.Data,
			}
			if got := s.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTask_Names(t *testing.T) {
	type fields struct {
		ID        string
		Name      string
		CreatedAt time.Time
		Status    string
		Data      string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success - task field names",
			fields: fields{
				ID:        "1",
				Name:      "Shahzad",
				CreatedAt: time.Now(),
				Status:    "active",
				Data:      "convert task struct to map",
			},
			want: []string{"id", "name", "created_at", "status", "data"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Task{
				ID:        tt.fields.ID,
				Name:      tt.fields.Name,
				CreatedAt: tt.fields.CreatedAt,
				Status:    tt.fields.Status,
				Data:      tt.fields.Data,
			}
			if got := s.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
