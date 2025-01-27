package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          int       `json:"id" db:"id"`
	UUID        uuid.UUID `json:"uuid" db:"uuid"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      int       `json:"status" db:"status"`
	UserID      int       `json:"user_id" db:"user_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

type TaskCreateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	UserID      int    `json:"user_id"`
}

type TaskUpdateRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type TaskUpdateStatusRequest struct {
	Status int `json:"status"`
}
