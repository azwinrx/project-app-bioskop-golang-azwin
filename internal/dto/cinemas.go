package dto

import (
	"time"
)

type CinemasRequest struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Address string `json:"address"`
}

type CinemasResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	City      string    `json:"city"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
