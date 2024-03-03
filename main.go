package main

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/controllers"
	"github.com/afanasyevadina/go-test/util"
	"net/http"
)

func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Accept", "application/json")
		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return apiMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := config.Authenticate(r); err != nil {
			util.JsonResponse(w, util.Message{
				Status:  http.StatusUnauthorized,
				Message: http.StatusText(http.StatusUnauthorized),
			}, http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}))
}

func main() {
	config.ConnectDB()
	taskController := controllers.NewTaskController()
	mux := http.NewServeMux()
	mux.Handle("/api/tasks", authMiddleware(http.HandlerFunc(taskController.TasksList)))
	http.ListenAndServe(":8080", mux)
}
