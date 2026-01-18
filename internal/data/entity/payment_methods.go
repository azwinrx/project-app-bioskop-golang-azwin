package entity

import "time"

type PaymentMethodsRepository struct {
	Id        int
	Name      string
	LogoUrl   *string
	IsActive  bool
	CreatedAt time.Time
}
