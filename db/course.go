package db

import (
	"time"
)

// The courses data struct
type Course struct {
	Id         *int       `json:"id"`
	CourseName *string    `json:"courseName" validate:"required"`
	City       *string    `json:"city" validate:"required"`
	State      *string    `json:"state" validate:"required"`
	Zip        *string    `json:"zip" validate:"required"`
	CreatedAt  *time.Time `json:"createdAt"`
	UpdatedAt  *time.Time `json:"updatedAt"`
	DeletedAt  *time.Time `json:"deletedAt"`
}

// Get course by id
func (c Course) Find(id int) Course {
	Conn.First(&c, id)
	return c
}

func (c Course) FindAll() []Course {
	var cs []Course
	Conn.Find(&cs)
	return cs
}

func (c Course) Update() {
	Conn.Save(&c)
}

// Get course by course name
func (c Course) ByCoursename(n string) Course {
	Conn.Where("courseName = ?", n).First(&c)
	return c
}
