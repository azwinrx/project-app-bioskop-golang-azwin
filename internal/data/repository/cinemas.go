package repository

import (
	"context"
	"fmt"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)

type CinemasRepository interface {
	GetAllCinemas(limit, offset int) ([]*entity.CinemasRepository, error)
	GetCinemaByID(cinemaID int) (*entity.CinemasRepository, error)
	CountCinemas() (int, error)
}

type cinemasRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewCinemasRepository(db database.PgxIface, log *zap.Logger) CinemasRepository {
	return &cinemasRepository{db: db, logger: log}
}

func (r *cinemasRepository) GetAllCinemas(limit, offset int) ([]*entity.CinemasRepository, error) {
	query := `
		SELECT id, name, city, address, created_at, updated_at, deleted_at
		FROM cinemas
		WHERE deleted_at IS NULL
		ORDER BY name
		LIMIT $1 OFFSET $2
	`
	
	rows, err := r.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		r.logger.Error("failed to get cinemas", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var cinemas []*entity.CinemasRepository
	for rows.Next() {
		cinema := &entity.CinemasRepository{}
		err := rows.Scan(
			&cinema.Id,
			&cinema.Name,
			&cinema.City,
			&cinema.Address,
			&cinema.CreatedAt,
			&cinema.UpdatedAt,
			&cinema.DeletedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan cinema", zap.Error(err))
			return nil, err
		}
		cinemas = append(cinemas, cinema)
	}

	return cinemas, nil
}

func (r *cinemasRepository) GetCinemaByID(cinemaID int) (*entity.CinemasRepository, error) {
	query := `
		SELECT id, name, city, address, created_at, updated_at, deleted_at
		FROM cinemas
		WHERE id = $1 AND deleted_at IS NULL
	`
	
	cinema := &entity.CinemasRepository{}
	err := r.db.QueryRow(context.Background(), query, cinemaID).Scan(
		&cinema.Id,
		&cinema.Name,
		&cinema.City,
		&cinema.Address,
		&cinema.CreatedAt,
		&cinema.UpdatedAt,
		&cinema.DeletedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get cinema by ID", 
			zap.Int("cinema_id", cinemaID),
			zap.Error(err),
		)
		return nil, fmt.Errorf("cinema not found")
	}
	
	return cinema, nil
}

func (r *cinemasRepository) CountCinemas() (int, error) {
	query := `SELECT COUNT(*) FROM cinemas WHERE deleted_at IS NULL`
	
	var count int
	err := r.db.QueryRow(context.Background(), query).Scan(&count)
	if err != nil {
		r.logger.Error("failed to count cinemas", zap.Error(err))
		return 0, err
	}
	
	return count, nil
}
