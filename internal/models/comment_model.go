package models

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        int       `json:"id" db:"id"`
	UUID      uuid.UUID `json:"uuid" db:"uuid"`
	Content   string    `json:"content" db:"content"`
	UserID    int       `json:"user_id" db:"user_id"`
	TaskID    int       `json:"task_id" db:"task_id"`
	ParentID  *int      `json:"parent_id" db:"parent_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type CommentCreateRequest struct {
	Content  string `json:"content"`
	UserID   int    `json:"user_id"`
	TaskID   int    `json:"task_id"`
	ParentID *int   `json:"parent_id"`
}

type CommentUpdateRequest struct {
	Content  string `json:"content"`
	ParentID *int   `json:"parent_id"`
}
