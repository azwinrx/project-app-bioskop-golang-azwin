package repository

import (
	"time"
)

type UsersRepository struct {
	Id int
	Username string
	Email string
	Password string
	IsVerified bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}