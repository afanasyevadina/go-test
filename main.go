package main

import (
	"github.com/afanasyevadina/go-test/config"
	"github.com/afanasyevadina/go-test/controllers"
	"github.com/afanasyevadina/go-test/dto"
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
			dto.RespondWith401(w)
			return
		}
		next.ServeHTTP(w, r)
	}))
}

func main() {
	config.ConnectDB()
	taskController := controllers.NewTaskController()
	authController := controllers.NewAuthController()
	profileController := controllers.NewProfileController()
	mux := http.NewServeMux()
	mux.Handle("/api/login", apiMiddleware(http.HandlerFunc(authController.Login)))
	mux.Handle("/api/register", apiMiddleware(http.HandlerFunc(authController.Register)))
	mux.Handle("/api/profile", authMiddleware(http.HandlerFunc(profileController.Show)))
	mux.Handle("/api/tasks", authMiddleware(http.HandlerFunc(taskController.TasksList)))
	http.ListenAndServe(":8080", mux)
}
