package models

import (
	"github.com/fatih/structs"
	"time"
)

// Student holds information for a student
type Task struct {
	ID        string    `json:"id" structs:"id"  bson:"_id" db:"id"`
	Name      string    `json:"name" structs:"name"  bson:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" structs:"created_at" bson:"created_at" db:"created_at"`
	Status    string    `json:"status" structs:"status" bson:"status" db:"status"`
	Data      string    `json:"data" structs:"data" bson:"data" db:"data"`
}

// Map converts structs to a map representation
func (s *Task) Map() map[string]interface{} {
	return structs.Map(s)
}

// Names returns the field names of Student model
func (s *Task) Names() []string {
	fields := structs.Fields(s)
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
