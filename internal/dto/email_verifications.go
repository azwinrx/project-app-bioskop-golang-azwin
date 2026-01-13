package dto

import (
	"time"
)

type EmailVerificationsRequest struct {
	Email   string `json:"email"`
	OtpCode string `json:"otp_code"`
}

type EmailVerificationsResponse struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	OtpCode   string    `json:"otp_code"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
}
