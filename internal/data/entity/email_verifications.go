package repository

import (
	"time"
)

type EmailVerificationsRepository struct {
	Id int
	Email string
	OtpCode string
	ExpiredAt time.Time
	CreatedAt time.Time
}