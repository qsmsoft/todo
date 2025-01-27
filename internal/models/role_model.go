package models

import "time"

type Role struct {
	ID        int       `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RoleCreateRequest struct {
	Name string `json:"name"`
}

type RoleUpdateRequest struct {
	Name string `json:"name"`
}
