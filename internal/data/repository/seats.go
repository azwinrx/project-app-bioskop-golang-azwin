package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"
	"time"

	"go.uber.org/zap"
)

type SeatsRepository interface {
	GetSeatsByCinemaAndShowtime(cinemaID int, showDate, showTime time.Time) ([]*entity.SeatsRepository, error)
	GetSeatByID(seatID int) (*entity.SeatsRepository, error)
}

type seatsRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewSeatsRepository(db database.PgxIface, log *zap.Logger) SeatsRepository {
	return &seatsRepository{db: db, logger: log}
}

func (r *seatsRepository) GetSeatsByCinemaAndShowtime(cinemaID int, showDate, showTime time.Time) ([]*entity.SeatsRepository, error) {
	query := `
		SELECT s.id, s.cinema_id, s.seat_number, s.created_at
		FROM seats s
		WHERE s.cinema_id = $1
		ORDER BY s.seat_number
	`
	
	rows, err := r.db.Query(context.Background(), query, cinemaID)
	if err != nil {
		r.logger.Error("failed to get seats", 
			zap.Int("cinema_id", cinemaID),
			zap.Error(err),
		)
		return nil, err
	}
	defer rows.Close()

	var seats []*entity.SeatsRepository
	for rows.Next() {
		seat := &entity.SeatsRepository{}
		err := rows.Scan(
			&seat.Id,
			&seat.CinemaId,
			&seat.SeatNumber,
			&seat.CreatedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan seat", zap.Error(err))
			return nil, err
		}
		seats = append(seats, seat)
	}

	return seats, nil
}

func (r *seatsRepository) GetSeatByID(seatID int) (*entity.SeatsRepository, error) {
	query := `
		SELECT id, cinema_id, seat_number, created_at
		FROM seats
		WHERE id = $1
	`
	
	seat := &entity.SeatsRepository{}
	err := r.db.QueryRow(context.Background(), query, seatID).Scan(
		&seat.Id,
		&seat.CinemaId,
		&seat.SeatNumber,
		&seat.CreatedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get seat by ID", 
			zap.Int("seat_id", seatID),
			zap.Error(err),
		)
		return nil, err
	}
	
	return seat, nil
}
