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
	Title  string     `json:"title"`
	Status TaskStatus `json:"status"`
	UserID int        `json:"-"`
	User   User       `json:"user"`
}
