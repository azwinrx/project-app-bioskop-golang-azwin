package dto

import (
	"time"
)

type PaymentMethodsRequest struct {
	Name     string `json:"name"`
	LogoURL  string `json:"logo_url"`
	IsActive bool   `json:"is_active"`
}

type PaymentMethodsResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	LogoURL   string    `json:"logo_url"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}
