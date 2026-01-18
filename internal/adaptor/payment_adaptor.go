package adaptor

import (
	"encoding/json"
	"net/http"
	"project-app-bioskop-golang-azwin/internal/dto"
	"project-app-bioskop-golang-azwin/internal/usecase"
	"project-app-bioskop-golang-azwin/pkg/utils"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type PaymentAdaptor struct {
	paymentUsecase usecase.PaymentUsecase
	validator      *validator.Validate
	logger         *zap.Logger
}

func NewPaymentAdaptor(paymentUsecase usecase.PaymentUsecase, validator *validator.Validate, logger *zap.Logger) *PaymentAdaptor {
	return &PaymentAdaptor{
		paymentUsecase: paymentUsecase,
		validator:      validator,
		logger:         logger,
	}
}

func (h *PaymentAdaptor) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {
	// Call usecase
	paymentMethods, err := h.paymentUsecase.GetPaymentMethods()
	if err != nil {
		h.logger.Error("failed to get payment methods", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to get payment methods", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "payment methods retrieved successfully", paymentMethods)
}

func (h *PaymentAdaptor) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req dto.PaymentRequest
	
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
	response, err := h.paymentUsecase.ProcessPayment(req)
	if err != nil {
		h.logger.Error("failed to process payment", zap.Error(err))
		utils.ResponseBadRequest(w, http.StatusInternalServerError, "failed to process payment", err.Error())
		return
	}

	utils.ResponseSuccess(w, http.StatusOK, "payment processed successfully", response)
}
