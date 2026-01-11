package repository

import (
	"time"
)

type BookingsRepository struct {
	Id int
	UserId int
	ShowtimeId int
	TotalPrice float64
	Status string
	BookinCode string
	CreatedAt time.Time
	UpdatedAt time.Time
}