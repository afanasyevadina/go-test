package models

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	StatusOpen       TaskStatus = "Open"
	StatusInProgress TaskStatus = "In Progress"
	StatusInReview   TaskStatus = "In Review"
	StatusDone       TaskStatus = "Done"
)

type Task struct {
	gorm.Model
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	AuthorID    int        `json:"-"`
	Author      User       `json:"author"`
	AssigneeID  int        `json:"-"`
	Assignee    User       `json:"assignee"`
}
