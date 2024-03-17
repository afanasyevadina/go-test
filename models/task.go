package models

import (
	"gorm.io/gorm"
)

type TaskStatus string

const (
	TaskStatusOpen       TaskStatus = "Open"
	TaskStatusInProgress TaskStatus = "In Progress"
	TaskStatusInReview   TaskStatus = "In Review"
	TaskStatusDone       TaskStatus = "Done"
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
