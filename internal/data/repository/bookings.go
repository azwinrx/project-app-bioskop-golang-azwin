package repository

import (
	"context"
	"fmt"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)

type BookingsRepository interface {
	CreateBooking(booking *entity.BookingsRepository) error
	GetBookingByID(bookingID int) (*entity.BookingsRepository, error)
	GetBookingsByUserID(userID int) ([]*entity.BookingsRepository, error)
	UpdateBookingStatus(bookingID int, status string) error
	CheckSeatBooked(showtimeID, seatID int) (bool, error)
}

type bookingsRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewBookingsRepository(db database.PgxIface, log *zap.Logger) BookingsRepository {
	return &bookingsRepository{db: db, logger: log}
}

func (r *bookingsRepository) CreateBooking(booking *entity.BookingsRepository) error {
	query := `
		INSERT INTO bookings (user_id, showtime_id, total_price, status, booking_code)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at
	`
	
	err := r.db.QueryRow(
		context.Background(), 
		query, 
		booking.UserId, 
		booking.ShowtimeId, 
		booking.TotalPrice, 
		booking.Status,
		booking.BookinCode,
	).Scan(&booking.Id, &booking.CreatedAt)
	
	if err != nil {
		r.logger.Error("failed to create booking", 
			zap.Int("user_id", booking.UserId),
			zap.Error(err),
		)
		return err
	}
	
	r.logger.Info("booking created successfully", 
		zap.Int("booking_id", booking.Id),
		zap.Int("user_id", booking.UserId),
	)
	
	return nil
}

func (r *bookingsRepository) GetBookingByID(bookingID int) (*entity.BookingsRepository, error) {
	query := `
		SELECT id, user_id, showtime_id, total_price, status, booking_code, created_at, updated_at
		FROM bookings
		WHERE id = $1
	`
	
	booking := &entity.BookingsRepository{}
	err := r.db.QueryRow(context.Background(), query, bookingID).Scan(
		&booking.Id,
		&booking.UserId,
		&booking.ShowtimeId,
		&booking.TotalPrice,
		&booking.Status,
		&booking.BookinCode,
		&booking.CreatedAt,
		&booking.UpdatedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get booking by ID", 
			zap.Int("booking_id", bookingID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("booking not found")
	}
	
	return booking, nil
}

func (r *bookingsRepository) GetBookingsByUserID(userID int) ([]*entity.BookingsRepository, error) {
	query := `
		SELECT id, user_id, showtime_id, total_price, status, booking_code, created_at, updated_at
		FROM bookings
		WHERE user_id = $1
		ORDER BY created_at DESC
	`
	
	rows, err := r.db.Query(context.Background(), query, userID)
	if err != nil {
		r.logger.Error("failed to get bookings by user ID", 
			zap.Int("user_id", userID),
			zap.Error(err),
		)
		return nil, err
	}
	defer rows.Close()

	var bookings []*entity.BookingsRepository
	for rows.Next() {
		booking := &entity.BookingsRepository{}
		err := rows.Scan(
			&booking.Id,
			&booking.UserId,
			&booking.ShowtimeId,
			&booking.TotalPrice,
			&booking.Status,
			&booking.BookinCode,
			&booking.CreatedAt,
			&booking.UpdatedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan booking", zap.Error(err))
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *bookingsRepository) UpdateBookingStatus(bookingID int, status string) error {
	query := `
		UPDATE bookings
		SET status = $1, updated_at = NOW()
		WHERE id = $2
	`
	
	_, err := r.db.Exec(context.Background(), query, status, bookingID)
	if err != nil {
		r.logger.Error("failed to update booking status", 
			zap.Int("booking_id", bookingID),
			zap.String("status", status),
			zap.Error(err),
		)
		return err
	}
	
	return nil
}

func (r *bookingsRepository) CheckSeatBooked(showtimeID, seatID int) (bool, error) {
	query := `
		SELECT EXISTS(
			SELECT 1 
			FROM booking_seats bs
			INNER JOIN bookings b ON bs.booking_id = b.id
			WHERE b.showtime_id = $1 AND bs.seat_id = $2 AND b.status != 'cancelled'
		)
	`
	
	var exists bool
	err := r.db.QueryRow(context.Background(), query, showtimeID, seatID).Scan(&exists)
	if err != nil {
		r.logger.Error("failed to check seat booked", 
			zap.Int("showtime_id", showtimeID),
			zap.Int("seat_id", seatID),
			zap.Error(err),
		)
		return false, err
	}
	
	return exists, nil
}
