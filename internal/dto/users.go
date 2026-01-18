package dto

import (
	"time"

	"github.com/google/uuid"
)

// User Registration and Authentication DTOs
type RegisterRequest struct {
Username string `json:"username" validate:"required,min=3,max=50"`
Email    string `json:"email" validate:"required,email"`
Password string `json:"password" validate:"required,min=6"`
}

type LoginRequest struct {
Email    string `json:"email" validate:"required,email"`
Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
ID       int    `json:"id"`
Username string `json:"username"`
Email    string `json:"email"`
}

type LoginResponse struct {
	Token     uuid.UUID `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
	User      AuthResponse `json:"user"`
}

// Legacy DTOs
type UsersRequest struct {
Username string
Email    string
Password string
}

type UsersResponse struct {
Id        int
Username  string
Email     string
CreatedAt time.Time
UpdatedAt time.Time
DeletedAt *time.Time
}
