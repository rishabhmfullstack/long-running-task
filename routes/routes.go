package routes

import (
	"net/http"
	"task-api/handlers"
	"task-api/middleware"
)

func SetupRoutes() {
	http.HandleFunc("/task", middleware.Authenticate(handlers.HandleTask))
	http.HandleFunc("/task/status/", middleware.Authenticate(handlers.HandleTaskStatus))
	http.HandleFunc("/task/output/", middleware.Authenticate(handlers.HandleTaskOutput))
}
