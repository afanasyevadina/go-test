package dto

import "github.com/afanasyevadina/go-test/models"

type TasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

type TaskResponse struct {
	ID     uint              `json:"id"`
	Title  string            `json:"title"`
	Status models.TaskStatus `json:"status"`
	User   taskUserResponse  `json:"user"`
}

type taskUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TaskRequest struct {
	Title string `json:"title"`
}

func TaskResponseFromModel(task models.Task) TaskResponse {
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

func TasksResponseFromModels(tasks []models.Task) TasksResponse {
	response := TasksResponse{
		Tasks: make([]TaskResponse, 0, len(tasks)),
	}
	for _, t := range tasks {
		response.Tasks = append(response.Tasks, TaskResponseFromModel(t))
	}
	return response
}
