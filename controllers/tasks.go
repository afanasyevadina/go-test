package controllers

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/models"
	"github.com/afanasyevadina/go-test/util"
	"net/http"
)

func TasksList(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		StoreTask(w, req)
	} else {
		tasks := make([]models.Task, 0)
		config.DB.Find(&tasks)
		util.JsonResponse(w, tasks)
	}
}

func StoreTask(w http.ResponseWriter, req *http.Request) {
	task := models.Task{}
	util.JsonRequest(w, req, &task)
	task.Status = models.StatusOpen
	config.DB.Create(&task)
	util.JsonResponse(w, task)
}
