package models

import (
	"time"
)

// The base interface used for each data model
type Model interface {
	TableName() string
	Find(id int) Model
	Save() Model
	Delete(id int) Model
}

// Gets and formats a stamp to be used for CreatedAt, UpdatedAt, and DeletedAt
func Stamp() *time.Time {
	t := time.Now()
	return &t
}
