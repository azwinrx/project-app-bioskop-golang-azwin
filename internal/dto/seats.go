package dto

import (
	"time"
)

type SeatsRequest struct {
	CinemaId   int    `json:"cinema_id"`
	SeatNumber string `json:"seat_number"`
}

type SeatsResponse struct {
	Id         int       `json:"id"`
	CinemaId   int       `json:"cinema_id"`
	SeatNumber string    `json:"seat_number"`
	CreatedAt  time.Time `json:"created_at"`
}
