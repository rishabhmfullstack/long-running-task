package main

import (
	"log"
	"net/http"

	"task-api/database"
	"task-api/routes"
)

func main() {

	db, err := database.InitDB("task.db")
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	routes.SetupRoutes()

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
