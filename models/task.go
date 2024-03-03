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

type TaskResponse struct {
	ID     uint             `json:"id"`
	Title  string           `json:"title"`
	Status TaskStatus       `json:"status"`
	User   taskUserResponse `json:"user"`
}

type taskUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (task Task) ToResponse() TaskResponse {
	return TaskResponse{
		ID:     task.ID,
		Title:  task.Title,
		Status: task.Status,
		User: taskUserResponse{
			ID:   task.User.ID,
			Name: task.User.Name,
		},
	}
}
