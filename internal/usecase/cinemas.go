package usecase

import (
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/internal/dto"

	"go.uber.org/zap"
)

type CinemaUsecase interface {
	GetAllCinemas(page, limit int) (*dto.CinemaListResponse, error)
	GetCinemaByID(cinemaID int) (*dto.CinemaResponse, error)
}

type cinemaUsecase struct {
	cinemaRepo repository.CinemasRepository
	logger     *zap.Logger
}

func NewCinemaUsecase(cinemaRepo repository.CinemasRepository, logger *zap.Logger) CinemaUsecase {
	return &cinemaUsecase{
		cinemaRepo: cinemaRepo,
		logger:     logger,
	}
}

func (u *cinemaUsecase) GetAllCinemas(page, limit int) (*dto.CinemaListResponse, error) {
	offset := (page - 1) * limit

	cinemas, err := u.cinemaRepo.GetAllCinemas(limit, offset)
	if err != nil {
		u.logger.Error("failed to get cinemas", zap.Error(err))
		return nil, err
	}

	total, err := u.cinemaRepo.CountCinemas()
	if err != nil {
		u.logger.Error("failed to count cinemas", zap.Error(err))
		return nil, err
	}

	var cinemaResponses []dto.CinemaResponse
	for _, cinema := range cinemas {
		cinemaResponses = append(cinemaResponses, dto.CinemaResponse{
			ID:      cinema.Id,
			Name:    cinema.Name,
			City:    cinema.City,
			Address: cinema.Address,
		})
	}

	totalPages := (total + limit - 1) / limit

	return &dto.CinemaListResponse{
		Cinemas: cinemaResponses,
		Pagination: dto.Pagination{
			Page:       page,
			Limit:      limit,
			TotalItems: total,
			TotalPages: totalPages,
		},
	}, nil
}

func (u *cinemaUsecase) GetCinemaByID(cinemaID int) (*dto.CinemaResponse, error) {
	cinema, err := u.cinemaRepo.GetCinemaByID(cinemaID)
	if err != nil {
		u.logger.Error("failed to get cinema by ID", zap.Int("cinema_id", cinemaID), zap.Error(err))
		return nil, err
	}

	return &dto.CinemaResponse{
		ID:      cinema.Id,
		Name:    cinema.Name,
		City:    cinema.City,
		Address: cinema.Address,
	}, nil
}
