package models

import "time"

type Task struct {
	ID        int       `json:"id"`
	Status    string    `json:"status"`
	Result    string    `json:"result,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt string    `json:"deleted_at,omitempty"`
}
