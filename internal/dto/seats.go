package dto

// Seat DTOs
type SeatResponse struct {
ID         int    `json:"id"`
SeatNumber string `json:"seat_number"`
IsBooked   bool   `json:"is_booked"`
}

type SeatsAvailabilityResponse struct {
CinemaID   int            `json:"cinema_id"`
CinemaName string         `json:"cinema_name"`
ShowDate   string         `json:"show_date"`
ShowTime   string         `json:"show_time"`
Seats      []SeatResponse `json:"seats"`
}
