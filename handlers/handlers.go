package handlers

import (
	"encoding/json"
	"log"

	"net/http"
	"strconv"
	"task-api/database"
	"task-api/models"
	"time"
)

func HandleTask(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		startTask(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTaskStatus(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskStatus(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func HandleTaskOutput(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTaskOutput(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func startTask(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	res, err := db.Exec("INSERT INTO tasks (status, result) VALUES (?, ?)", "pending", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	go performLongRunningTask(int(id))

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]int{"task_id": int(id)})
}

func getTaskStatus(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Path[len("/task/status/"):]
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	var task models.Task
	err = db.QueryRow("SELECT id, status, result FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Status, &task.Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(task)
}

func getTaskOutput(w http.ResponseWriter, r *http.Request) {
	taskID := r.URL.Path[len("/task/output/"):]
	id, err := strconv.Atoi(taskID)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	db := database.GetDB()
	var task models.Task
	err = db.QueryRow("SELECT id, status, result FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Status, &task.Result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if task.Status != "completed" {
		http.Error(w, "Task not completed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"output": task.Result})
}

func performLongRunningTask(id int) {
	time.Sleep(10 * time.Second)

	db := database.GetDB()

	_, err := db.Exec(`
		UPDATE tasks SET status = ?, result = ?, updated_at = CURRENT_TIMESTAMP 
		WHERE id = ?`, "completed", "Task completed successfully", id)
	if err != nil {
		log.Printf("Failed to update task %d: %v", id, err)
	}
}
