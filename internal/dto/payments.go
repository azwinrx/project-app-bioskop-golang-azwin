package dto

// Payment Method DTOs
type PaymentMethodResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	LogoURL  string `json:"logo_url"`
	IsActive bool   `json:"is_active"`
}

// Payment DTOs
type PaymentRequest struct {
BookingID       int `json:"booking_id" validate:"required"`
PaymentMethodID int `json:"payment_method_id" validate:"required"`
}

type PaymentResponse struct {
PaymentID       int     `json:"payment_id"`
BookingID       int     `json:"booking_id"`
Amount          float64 `json:"amount"`
PaymentMethod   string  `json:"payment_method"`
Status          string  `json:"status"`
TransactionTime string  `json:"transaction_time"`
}
