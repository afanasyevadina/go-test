package controllers

import (
	"github.com/afanasyevadina/go-test/dto"
	"github.com/afanasyevadina/go-test/repositories"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func NewTaskController() TaskController {
	return TaskController{
		taskRepository: repositories.NewTaskRepository(),
		validator:      validator.New(validator.WithRequiredStructEnabled()),
	}
}

type TaskController struct {
	taskRepository *repositories.TaskRepository
	validator      *validator.Validate
}

func (c *TaskController) TasksList(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		c.StoreTask(w, r)
	} else {
		tasks := c.taskRepository.GetAll()
		dto.ToJsonResponse(w, dto.TasksResponseFromModels(tasks).Tasks, http.StatusOK)
	}
}

func (c *TaskController) StoreTask(w http.ResponseWriter, r *http.Request) {
	taskRequest := dto.TaskCreateRequest{}
	dto.FromRequest(r, &taskRequest)
	if err := c.validator.Struct(taskRequest); err != nil {
		dto.ToJsonResponse(w, dto.ResponseFromValidator(err.(validator.ValidationErrors)), http.StatusUnprocessableEntity)
		return
	}
	task := c.taskRepository.Create(taskRequest.ToModel())
	dto.ToJsonResponse(w, dto.TaskResponseFromModel(task), http.StatusCreated)
}
