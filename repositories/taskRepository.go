package repositories

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
)

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

type TaskRepository struct {
}

func (repo TaskRepository) GetAll() []models.Task {
	tasks := make([]models.Task, 0)
	config.DB.Preload("Author").Preload("Assignee").Find(&tasks)
	return tasks
}

func (repo TaskRepository) Create(task models.Task) models.Task {
	task.Status = models.TaskStatusOpen
	task.Author = *config.CurrentUser
	config.DB.Save(&task)
	return task
}
