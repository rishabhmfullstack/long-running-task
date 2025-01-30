package tasks

import (
	"task-api/database"
	"time"
)

func PerformLongRunningTask(id int) {
	time.Sleep(10 * time.Second)
	database.GetDB().Exec("UPDATE tasks SET status = ?, result = ? WHERE id = ?", "completed", "Task completed successfully", id)
}
