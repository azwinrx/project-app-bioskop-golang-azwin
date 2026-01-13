package dto

import (
	"time"
)

type BookingsRequest struct {
	UserId      int     `json:"user_id"`
	ShowtimeId  int     `json:"showtime_id"`
	TotalPrice  float64 `json:"total_price"`
	Status      string  `json:"status"`
	BookingCode string  `json:"booking_code"`
}

type BookingsResponse struct {
	Id          int       `json:"id"`
	UserId      int       `json:"user_id"`
	ShowtimeId  int       `json:"showtime_id"`
	TotalPrice  float64   `json:"total_price"`
	Status      string    `json:"status"`
	BookingCode string    `json:"booking_code"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
