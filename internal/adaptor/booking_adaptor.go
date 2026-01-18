package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-azwin/internal/dto"
	"project-app-bioskop-golang-azwin/internal/middleware"
	"project-app-bioskop-golang-azwin/internal/usecase"
	"project-app-bioskop-golang-azwin/pkg/utils"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type BookingAdaptor struct {
	bookingUsecase usecase.BookingUsecase
	validator      *validator.Validate
	logger         *zap.Logger
}

func NewBookingAdaptor(bookingUsecase usecase.BookingUsecase, validator *validator.Validate, logger *zap.Logger) *BookingAdaptor {
	return &BookingAdaptor{
		bookingUsecase: bookingUsecase,
		validator:      validator,
		logger:         logger,
	}
}

func (h *BookingAdaptor) GetSeatsAvailability(w http.ResponseWriter, r *http.Request) {
	cinemaIDStr := chi.URLParam(r, "cinemaId")
	cinemaID, err := strconv.Atoi(cinemaIDStr)
	if err != nil {
		h.logger.Warn("invalid cinema ID", zap.String("cinema_id", cinemaIDStr))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid cinema ID", "cinema ID must be a number")
		return
	}

	date := r.URL.Query().Get("date")
	time := r.URL.Query().Get("time")

	if date == "" || time == "" {
		h.logger.Warn("missing date or time parameters")
		utils.ResponseBadRequest(w, http.StatusBadRequest, "missing parameters", "date and time are required")
		return
	}

	// Call usecase
	response, err := h.bookingUsecase.GetSeatsAvailability(cinemaID, date, time)
	if err != nil {
		h.logger.Error("failed to get seats availability", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to get seats availability", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "seats availability retrieved successfully", response)
}

func (h *BookingAdaptor) CreateBooking(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		h.logger.Error("failed to get user from context")
		utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "user not found in context")
		return
	}

	var req dto.BookingRequest
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("failed to decode request", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "invalid request body", err.Error())
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		h.logger.Warn("validation failed", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusBadRequest, "validation failed", err.Error())
		return
	}

	// Call usecase
	response, err := h.bookingUsecase.CreateBooking(req, claims.UserID)
	if err != nil {
		h.logger.Error("failed to create booking", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to create booking", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "booking created successfully", response)
}

func (h *BookingAdaptor) GetUserBookings(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		h.logger.Error("failed to get user from context")
		utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "user not found in context")
		return
	}

	// Call usecase
	bookings, err := h.bookingUsecase.GetUserBookings(claims.UserID)
	if err != nil {
		h.logger.Error("failed to get user bookings", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to get user bookings", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "user bookings retrieved successfully", bookings)
}
