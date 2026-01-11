package repository

import (
	"time"
)

type MoviesRepository struct {
	Id              int
	Title           string
	Description     string
	DurationMinutes int // duration in minutes
	Genre           string
	PosterURL       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time
}