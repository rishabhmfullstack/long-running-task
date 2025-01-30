package models

type Task struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
	Result string `json:"result,omitempty"`
}
