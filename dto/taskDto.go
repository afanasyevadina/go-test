package dto

import "github.com/afanasyevadina/go-test/models"

type TasksResponse struct {
	Tasks []TaskResponse `json:"tasks"`
}

type TaskResponse struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      models.TaskStatus `json:"status"`
	Author      taskUserResponse  `json:"author"`
	Assignee    taskUserResponse  `json:"assignee"`
}

type taskUserResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func TaskResponseFromModel(task models.Task) TaskResponse {
	return TaskResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		Author: taskUserResponse{
			ID:   task.Author.ID,
			Name: task.Author.Name,
		},
		Assignee: taskUserResponse{
			ID:   task.Assignee.ID,
			Name: task.Assignee.Name,
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
