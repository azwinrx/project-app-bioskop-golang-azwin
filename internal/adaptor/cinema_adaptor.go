package adaptor

import (
	"net/http"
	"project-app-bioskop-golang-azwin/internal/usecase"
	"project-app-bioskop-golang-azwin/pkg/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type CinemaAdaptor struct {
	cinemaUsecase usecase.CinemaUsecase
	logger        *zap.Logger
}

func NewCinemaAdaptor(cinemaUsecase usecase.CinemaUsecase, logger *zap.Logger) *CinemaAdaptor {
	return &CinemaAdaptor{
		cinemaUsecase: cinemaUsecase,
		logger:        logger,
	}
}

func (h *CinemaAdaptor) GetAllCinemas(w http.ResponseWriter, r *http.Request) {
	// Get pagination parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 10 // default limit

	if pageStr != "" {
		if p, err := strconv.Atoi(pageStr); err == nil && p > 0 {
			page = p
		}
	}

	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	// Call usecase
	response, err := h.cinemaUsecase.GetAllCinemas(page, limit)
	if err != nil {
		h.logger.Error("failed to get cinemas", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to get cinemas", err.Error())
		return
	}

	utils.ResponsePagination(w, http.StatusOK, "cinemas retrieved successfully", response.Cinemas, response.Pagination)
}

func (h *CinemaAdaptor) GetCinemaByID(w http.ResponseWriter, r *http.Request) {
	cinemaIDStr := chi.URLParam(r, "cinemaId")
	cinemaID, err := strconv.Atoi(cinemaIDStr)
	if err != nil {
		h.logger.Warn("invalid cinema ID", zap.String("cinema_id", cinemaIDStr))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid cinema ID", "cinema ID must be a number")
		return
	}

	// Call usecase
	cinema, err := h.cinemaUsecase.GetCinemaByID(cinemaID)
	if err != nil {
		h.logger.Error("failed to get cinema", zap.Int("cinema_id", cinemaID), zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusNotFound, "cinema not found", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "cinema retrieved successfully", cinema)
}
