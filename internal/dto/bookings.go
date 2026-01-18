package dto

// Booking DTOs
type BookingRequest struct {
CinemaID      int    `json:"cinema_id" validate:"required"`
SeatID        int    `json:"seat_id" validate:"required"`
Date          string `json:"date" validate:"required"`
Time          string `json:"time" validate:"required"`
PaymentMethod int    `json:"payment_method" validate:"required"`
}

type BookingResponse struct {
BookingID   int     `json:"booking_id"`
BookingCode string  `json:"booking_code"`
CinemaID    int     `json:"cinema_id"`
SeatNumber  string  `json:"seat_number"`
ShowDate    string  `json:"show_date"`
ShowTime    string  `json:"show_time"`
TotalPrice  float64 `json:"total_price"`
Status      string  `json:"status"`
}

type BookingHistoryResponse struct {
BookingID   int      `json:"booking_id"`
BookingCode string   `json:"booking_code"`
CinemaName  string   `json:"cinema_name"`
ShowDate    string   `json:"show_date"`
ShowTime    string   `json:"show_time"`
SeatNumbers []string `json:"seat_numbers"`
TotalPrice  float64  `json:"total_price"`
Status      string   `json:"status"`
BookingDate string   `json:"booking_date"`
}
