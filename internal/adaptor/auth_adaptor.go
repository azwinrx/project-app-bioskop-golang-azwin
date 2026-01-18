package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-azwin/internal/dto"
	"project-app-bioskop-golang-azwin/internal/middleware"
	"project-app-bioskop-golang-azwin/internal/usecase"
	"project-app-bioskop-golang-azwin/pkg/utils"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthAdaptor struct {
	authUsecase usecase.AuthUsecase
	validator   *validator.Validate
	logger      *zap.Logger
}

func NewAuthAdaptor(authUsecase usecase.AuthUsecase, validator *validator.Validate, logger *zap.Logger) *AuthAdaptor {
	return &AuthAdaptor{
		authUsecase: authUsecase,
		validator:   validator,
		logger:      logger,
	}
}

func (h *AuthAdaptor) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	
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
	user, err := h.authUsecase.Register(req)
	if err != nil {
		h.logger.Error("failed to register user", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to register user", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusCreated, "user registered successfully", user)
}

func (h *AuthAdaptor) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	
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
	response, err := h.authUsecase.Login(req)
	if err != nil {
		h.logger.Error("login failed", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusUnauthorized, "login failed", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "login successful", response)
}

func (h *AuthAdaptor) Logout(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	claims, ok := middleware.GetUserFromContext(r.Context())
	if !ok {
		h.logger.Error("failed to get user from context")
		utils.ResponseBadRequest(w, http.StatusUnauthorized, "unauthorized", "user not found in context")
		return
	}

	// Call usecase with session ID
	err := h.authUsecase.Logout(claims.SessionID)
	if err != nil {
		h.logger.Error("logout failed", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "logout failed", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "logout successful", nil)
}
