package dto

import (
	"time"
)

type PaymentsRequest struct {
	BookingId       int     `json:"booking_id"`
	PaymentMethodId int     `json:"payment_method_id"`
	Amount          float64 `json:"amount"`
	Status          string  `json:"status"`
}

type PaymentsResponse struct {
	Id              int       `json:"id"`
	BookingId       int       `json:"booking_id"`
	PaymentMethodId int       `json:"payment_method_id"`
	Amount          float64   `json:"amount"`
	TransactionTime time.Time `json:"transaction_time"`
	Status          string    `json:"status"`
}
