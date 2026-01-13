package dto

import (
	"time"
)

type MoviesRequest struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes"`
	Genre           string `json:"genre"`
	PosterURL       string `json:"poster_url"`
}

type MoviesResponse struct {
	Id              int       `json:"id"`
	Title           string    `json:"title"`
	Description     string    `json:"description"`
	DurationMinutes int       `json:"duration_minutes"`
	Genre           string    `json:"genre"`
	PosterURL       string    `json:"poster_url"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
