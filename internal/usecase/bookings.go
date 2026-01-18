package usecase

import (
	"errors"
	"fmt"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/internal/dto"
	"project-app-bioskop-golang-azwin/pkg/utils"
	"time"

	"go.uber.org/zap"
)

type BookingUsecase interface {
	GetSeatsAvailability(cinemaID int, date, timeStr string) (*dto.SeatsAvailabilityResponse, error)
	CreateBooking(req dto.BookingRequest, userID int) (*dto.BookingResponse, error)
	GetUserBookings(userID int) ([]dto.BookingHistoryResponse, error)
}

type bookingUsecase struct {
	bookingRepo      repository.BookingsRepository
	bookingSeatRepo  repository.BookingSeatsRepository
	seatRepo         repository.SeatsRepository
	showtimeRepo     repository.ShowtimesRepository
	cinemaRepo       repository.CinemasRepository
	logger           *zap.Logger
}

func NewBookingUsecase(
	bookingRepo repository.BookingsRepository,
	bookingSeatRepo repository.BookingSeatsRepository,
	seatRepo repository.SeatsRepository,
	showtimeRepo repository.ShowtimesRepository,
	cinemaRepo repository.CinemasRepository,
	logger *zap.Logger,
) BookingUsecase {
	return &bookingUsecase{
		bookingRepo:     bookingRepo,
		bookingSeatRepo: bookingSeatRepo,
		seatRepo:        seatRepo,
		showtimeRepo:    showtimeRepo,
		cinemaRepo:      cinemaRepo,
		logger:          logger,
	}
}

func (u *bookingUsecase) GetSeatsAvailability(cinemaID int, date, timeStr string) (*dto.SeatsAvailabilityResponse, error) {
	// Parse date and time
	showDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	showTime, err := time.Parse("15:04", timeStr)
	if err != nil {
		return nil, errors.New("invalid time format, use HH:MM")
	}

	// Get showtime
	showtime, err := u.showtimeRepo.GetShowtimeByCinemaAndDateTime(cinemaID, showDate, showTime)
	if err != nil {
		u.logger.Error("showtime not found", zap.Error(err))
		return nil, errors.New("showtime not found for the given cinema, date, and time")
	}

	// Get all seats for cinema
	seats, err := u.seatRepo.GetSeatsByCinemaAndShowtime(cinemaID, showDate, showTime)
	if err != nil {
		u.logger.Error("failed to get seats", zap.Error(err))
		return nil, err
	}

	// Check which seats are booked
	var seatResponses []dto.SeatResponse
	for _, seat := range seats {
		isBooked, _ := u.bookingRepo.CheckSeatBooked(showtime.Id, seat.Id)
		seatResponses = append(seatResponses, dto.SeatResponse{
			ID:         seat.Id,
			SeatNumber: seat.SeatNumber,
			IsBooked:   isBooked,
		})
	}

	// Get cinema info
	cinema, _ := u.cinemaRepo.GetCinemaByID(cinemaID)

	return &dto.SeatsAvailabilityResponse{
		CinemaID:   cinemaID,
		CinemaName: cinema.Name,
		ShowDate:   showDate.Format("2006-01-02"),
		ShowTime:   showTime.Format("15:04"),
		Seats:      seatResponses,
	}, nil
}

func (u *bookingUsecase) CreateBooking(req dto.BookingRequest, userID int) (*dto.BookingResponse, error) {
	// Parse date and time
	showDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return nil, errors.New("invalid date format, use YYYY-MM-DD")
	}

	showTime, err := time.Parse("15:04", req.Time)
	if err != nil {
		return nil, errors.New("invalid time format, use HH:MM")
	}

	// Get showtime
	showtime, err := u.showtimeRepo.GetShowtimeByCinemaAndDateTime(req.CinemaID, showDate, showTime)
	if err != nil {
		u.logger.Error("showtime not found", zap.Error(err))
		return nil, errors.New("showtime not found")
	}

	// Check if seat exists and belongs to the cinema
	seat, err := u.seatRepo.GetSeatByID(req.SeatID)
	if err != nil {
		return nil, errors.New("seat not found")
	}

	if seat.CinemaId != req.CinemaID {
		return nil, errors.New("seat does not belong to the selected cinema")
	}

	// Check if seat is already booked
	isBooked, err := u.bookingRepo.CheckSeatBooked(showtime.Id, req.SeatID)
	if err != nil {
		u.logger.Error("failed to check seat availability", zap.Error(err))
		return nil, errors.New("failed to check seat availability")
	}

	if isBooked {
		return nil, errors.New("seat is already booked")
	}

	// Generate booking code
	bookingCode := fmt.Sprintf("BK-%s", utils.GenerateUUIDToken()[:8])

	// Create booking
	booking := &entity.BookingsRepository{
		UserId:      userID,
		ShowtimeId:  showtime.Id,
		TotalPrice:  showtime.Price,
		Status:      "pending",
		BookinCode:  bookingCode,
	}

	err = u.bookingRepo.CreateBooking(booking)
	if err != nil {
		u.logger.Error("failed to create booking", zap.Error(err))
		return nil, errors.New("failed to create booking")
	}

	// Create booking seat
	bookingSeat := &entity.BookingSeatsRepository{
		BookingId: booking.Id,
		SeatId:    req.SeatID,
	}

	err = u.bookingSeatRepo.CreateBookingSeat(bookingSeat)
	if err != nil {
		u.logger.Error("failed to create booking seat", zap.Error(err))
		return nil, errors.New("failed to create booking seat")
	}

	return &dto.BookingResponse{
		BookingID:   booking.Id,
		BookingCode: booking.BookinCode,
		CinemaID:    req.CinemaID,
		SeatNumber:  seat.SeatNumber,
		ShowDate:    showDate.Format("2006-01-02"),
		ShowTime:    showTime.Format("15:04"),
		TotalPrice:  booking.TotalPrice,
		Status:      booking.Status,
	}, nil
}

func (u *bookingUsecase) GetUserBookings(userID int) ([]dto.BookingHistoryResponse, error) {
	bookings, err := u.bookingRepo.GetBookingsByUserID(userID)
	if err != nil {
		u.logger.Error("failed to get user bookings", zap.Int("user_id", userID), zap.Error(err))
		return nil, err
	}

	var bookingHistories []dto.BookingHistoryResponse
	for _, booking := range bookings {
		// Get showtime details
		showtime, err := u.showtimeRepo.GetShowtimeByID(booking.ShowtimeId)
		if err != nil {
			u.logger.Warn("failed to get showtime for booking", zap.Int("booking_id", booking.Id))
			continue
		}

		// Get cinema details
		cinema, err := u.cinemaRepo.GetCinemaByID(showtime.CinemaId)
		if err != nil {
			u.logger.Warn("failed to get cinema for booking", zap.Int("booking_id", booking.Id))
			continue
		}

		// Get booking seats
		bookingSeats, err := u.bookingSeatRepo.GetBookingSeatsByBookingID(booking.Id)
		if err != nil {
			u.logger.Warn("failed to get booking seats", zap.Int("booking_id", booking.Id))
			continue
		}

		var seatNumbers []string
		for _, bs := range bookingSeats {
			seat, err := u.seatRepo.GetSeatByID(bs.SeatId)
			if err == nil {
				seatNumbers = append(seatNumbers, seat.SeatNumber)
			}
		}

		bookingHistories = append(bookingHistories, dto.BookingHistoryResponse{
			BookingID:   booking.Id,
			BookingCode: booking.BookinCode,
			CinemaName:  cinema.Name,
			ShowDate:    showtime.ShowDate.Format("2006-01-02"),
			ShowTime:    showtime.ShowTime.Format("15:04"),
			SeatNumbers: seatNumbers,
			TotalPrice:  booking.TotalPrice,
			Status:      booking.Status,
			BookingDate: booking.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return bookingHistories, nil
}
