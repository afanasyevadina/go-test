package main

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/controllers"
	"github.com/afanasyevadina/go-test/util"
	"net/http"
)

func main() {
	config.ConnectDB()
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		util.JsonResponse(w, req.Header)
	})
	http.HandleFunc("/tasks", controllers.TasksList)
	http.ListenAndServe(":8080", nil)
}
