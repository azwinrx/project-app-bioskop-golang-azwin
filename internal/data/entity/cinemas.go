package entity

import (
	"time"
)

type CinemasRepository struct {
	Id int
	Name string
	City string
	Address string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}