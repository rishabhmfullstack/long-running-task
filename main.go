package main

import (
	"log"
	"net/http"

	"task-api/database"
	"task-api/handlers"
	"task-api/middleware"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.InitDB("task.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/task", middleware.Authenticate(handlers.HandleTask))
	http.HandleFunc("/task/status/", middleware.Authenticate(handlers.HandleTaskStatus))
	http.HandleFunc("/task/output/", middleware.Authenticate(handlers.HandleTaskOutput))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
