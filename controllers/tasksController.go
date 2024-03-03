package controllers

import (
	"encoding/json"
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
	"github.com/afanasyevadina/go-test/util"
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
		config.DB.Find(&tasks)
		responses := make([]interface{}, 0)
		for _, t := range tasks {
			responses = append(responses, t.ToResponse())
		}
		util.JsonResponse(w, responses, http.StatusOK)
	}
}

func (c *TaskController) StoreTask(w http.ResponseWriter, r *http.Request) {
	task := models.Task{}
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		util.JsonResponse(w, util.Message{
			Status:  http.StatusBadRequest,
			Message: http.StatusText(http.StatusBadRequest),
		}, http.StatusBadRequest)
		return
	}
	task.Status = models.StatusOpen
	user := models.User{}
	config.DB.Take(&user)
	task.User = user
	config.DB.Create(&task)
	util.JsonResponse(w, task.ToResponse(), http.StatusCreated)
}
