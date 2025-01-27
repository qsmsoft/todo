package models

type UserRole struct {
	UserID int `json:"user_id" db:"user_id"`
	RoleID int `json:"role_id" db:"role_id"`
}
