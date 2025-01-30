package models

import "time"

type Task struct {
	ID        int        `json:"id"`
	Status    string     `json:"status"`
	Output    string     `json:"output"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
