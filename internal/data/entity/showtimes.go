package entity

import (
	"time"
)

type ShowtimesRepository struct {
	Id        int
	MovieId   int
	CinemaId  int
	ShowDate time.Time
	ShowTime time.Time
	Price     float64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}