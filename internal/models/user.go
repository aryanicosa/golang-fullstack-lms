package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID           uuid.UUID `json:"id" validate:"required,uuid"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Email        string    `json:"email" validate:"required,email,lte=255"`
	Username     string    `json:"username" validate:"required,lte=255"`
	PasswordHash string    `json:"password_hash,omitempty" validate:"required,lte=255"`
	UserStatus   int       `json:"user_status" validate:"required,len=1"`
	UserRole     string    `json:"user_role" validate:"required,lte=25"`
	DisplayName  string    `json:"display_name" validate:"required,lte=255"`
}

type UserCreateRequest struct {
	Email    string `json:"email" validate:"required,email,lte=255"`
	Username string `json:"username" validate:"required,lte=255"`
	Password string `json:"password" validate:"required,lte=255"`
	DisplayName string `json:"display_name" validate:"required,lte=255"`
}

type UserCreateResponse struct {
	UserID uuid.UUID `json:"user_id" validate:"required,uuid"`
}