package repository

import "time"

type SeatsRepository struct {
	Id         int
	CinemaId   int
	SeatNumber string
	CreatedAt  time.Time
}