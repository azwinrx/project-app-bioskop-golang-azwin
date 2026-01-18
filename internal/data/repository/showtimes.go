package repository

import (
	"context"
	"fmt"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"
	"time"

	"go.uber.org/zap"
)

type ShowtimesRepository interface {
	GetShowtimeByCinemaAndDateTime(cinemaID int, showDate, showTime time.Time) (*entity.ShowtimesRepository, error)
	GetShowtimeByID(showtimeID int) (*entity.ShowtimesRepository, error)
}

type showtimesRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewShowtimesRepository(db database.PgxIface, log *zap.Logger) ShowtimesRepository {
	return &showtimesRepository{db: db, logger: log}
}

func (r *showtimesRepository) GetShowtimeByCinemaAndDateTime(cinemaID int, showDate, showTime time.Time) (*entity.ShowtimesRepository, error) {
	// Format show_time to extract only HH:MM:SS for comparison
	timeStr := showTime.Format("15:04:05")
	
	query := `
		SELECT id, movie_id, cinema_id, show_date, show_time, price, created_at, updated_at, deleted_at
		FROM showtimes
		WHERE cinema_id = $1 
		AND show_date = $2::date
		AND show_time::text = $3
		AND deleted_at IS NULL
	`
	
	showtime := &entity.ShowtimesRepository{}
	err := r.db.QueryRow(context.Background(), query, cinemaID, showDate, timeStr).Scan(
		&showtime.Id,
		&showtime.MovieId,
		&showtime.CinemaId,
		&showtime.ShowDate,
		&showtime.ShowTime,
		&showtime.Price,
		&showtime.CreatedAt,
		&showtime.UpdatedAt,
		&showtime.DeletedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get showtime", 
			zap.Int("cinema_id", cinemaID),
			zap.Time("show_date", showDate),
			zap.Time("show_time", showTime),
			zap.Error(err),
		)
		return nil, fmt.Errorf("showtime not found")
	}
	
	return showtime, nil
}

func (r *showtimesRepository) GetShowtimeByID(showtimeID int) (*entity.ShowtimesRepository, error) {
	query := `
		SELECT id, movie_id, cinema_id, show_date, show_time, price, created_at, updated_at, deleted_at
		FROM showtimes
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	showtime := &entity.ShowtimesRepository{}
	err := r.db.QueryRow(context.Background(), query, showtimeID).Scan(
		&showtime.Id,
		&showtime.MovieId,
		&showtime.CinemaId,
		&showtime.ShowDate,
		&showtime.ShowTime,
		&showtime.Price,
		&showtime.CreatedAt,
		&showtime.UpdatedAt,
		&showtime.DeletedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get showtime by ID", 
			zap.Int("showtime_id", showtimeID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("showtime not found")
	}
	
	return showtime, nil
}
