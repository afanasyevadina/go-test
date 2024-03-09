package controllers

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/dto"
	"github.com/afanasyevadina/go-test/models"
	"net/http"
)

func NewTaskController() TaskController {
	return TaskController{}
}

type TaskController struct {
}

func (c *TaskController) TasksList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.StoreTask(w, r)
	} else {
		tasks := make([]models.Task, 0)
		config.DB.Preload("User").Find(&tasks)
		dto.ToJsonResponse(w, dto.TasksResponseFromModels(tasks), http.StatusOK)
	}
}

func (c *TaskController) StoreTask(w http.ResponseWriter, r *http.Request) {
	taskRequest := dto.TaskRequest{}
	dto.FromRequest(r, &taskRequest)
	task := models.Task{Title: taskRequest.Title, Status: models.StatusOpen, User: *config.CurrentUser}
	config.DB.Create(&task)
	dto.ToJsonResponse(w, dto.TaskResponseFromModel(task), http.StatusCreated)
}
