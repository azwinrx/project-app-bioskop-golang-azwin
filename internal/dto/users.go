package dto

import (
	"time"
)
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