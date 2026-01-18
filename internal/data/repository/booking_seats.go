package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)

type BookingSeatsRepository interface {
	CreateBookingSeat(bookingSeat *entity.BookingSeatsRepository) error
	GetBookingSeatsByBookingID(bookingID int) ([]*entity.BookingSeatsRepository, error)
}

type bookingSeatsRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewBookingSeatsRepository(db database.PgxIface, log *zap.Logger) BookingSeatsRepository {
	return &bookingSeatsRepository{db: db, logger: log}
}

func (r *bookingSeatsRepository) CreateBookingSeat(bookingSeat *entity.BookingSeatsRepository) error {
	query := `
		INSERT INTO booking_seats (booking_id, seat_id)
		VALUES ($1, $2)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		context.Background(), 
		query, 
		bookingSeat.BookingId, 
		bookingSeat.SeatId,
	).Scan(&bookingSeat.Id)
	
	if err != nil {
		r.logger.Error("failed to create booking seat", 
			zap.Int("booking_id", bookingSeat.BookingId),
			zap.Int("seat_id", bookingSeat.SeatId),
			zap.Error(err),
		)
		return err
	}
	
	return nil
}

func (r *bookingSeatsRepository) GetBookingSeatsByBookingID(bookingID int) ([]*entity.BookingSeatsRepository, error) {
	query := `
		SELECT id, booking_id, seat_id
		FROM booking_seats
		WHERE booking_id = $1
	`
	
	rows, err := r.db.Query(context.Background(), query, bookingID)
	if err != nil {
		r.logger.Error("failed to get booking seats", 
			zap.Int("booking_id", bookingID),
			zap.Error(err),
		)
		return nil, err
	}
	defer rows.Close()

	var bookingSeats []*entity.BookingSeatsRepository
	for rows.Next() {
		bookingSeat := &entity.BookingSeatsRepository{}
		err := rows.Scan(
			&bookingSeat.Id,
			&bookingSeat.BookingId,
			&bookingSeat.SeatId,
		)
		if err != nil {
			r.logger.Error("failed to scan booking seat", zap.Error(err))
			return nil, err
		}
		bookingSeats = append(bookingSeats, bookingSeat)
	}

	return bookingSeats, nil
}
