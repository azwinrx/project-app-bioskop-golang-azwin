package entity

import (
	"time"
)

type PaymentsRepository struct {
	Id int
	BookingId int
	PaymentMethodId int
	Amount float64
	TransactionTime time.Time
	Status string
}