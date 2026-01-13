package dto

import (
	"time"
)

type ShowtimesRequest struct {
	MovieId  int       `json:"movie_id"`
	CinemaId int       `json:"cinema_id"`
	ShowDate time.Time `json:"show_date"`
	ShowTime time.Time `json:"show_time"`
	Price    float64   `json:"price"`
}

type ShowtimesResponse struct {
	Id        int       `json:"id"`
	MovieId   int       `json:"movie_id"`
	CinemaId  int       `json:"cinema_id"`
	ShowDate  time.Time `json:"show_date"`
	ShowTime  time.Time `json:"show_time"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
