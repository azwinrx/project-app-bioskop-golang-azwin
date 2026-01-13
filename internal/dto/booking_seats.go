package dto

type BookingSeatsRequest struct {
	BookingId int `json:"booking_id"`
	SeatId    int `json:"seat_id"`
}

type BookingSeatsResponse struct {
	Id        int `json:"id"`
	BookingId int `json:"booking_id"`
	SeatId    int `json:"seat_id"`
}
