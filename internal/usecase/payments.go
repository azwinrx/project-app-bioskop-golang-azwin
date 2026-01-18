package usecase

import (
	"errors"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/internal/dto"
	"time"

	"go.uber.org/zap"
)

type PaymentUsecase interface {
	GetPaymentMethods() ([]dto.PaymentMethodResponse, error)
	ProcessPayment(req dto.PaymentRequest) (*dto.PaymentResponse, error)
}

type paymentUsecase struct {
	paymentRepo       repository.PaymentsRepository
	paymentMethodRepo repository.PaymentMethodsRepository
	bookingRepo       repository.BookingsRepository
	logger            *zap.Logger
}

func NewPaymentUsecase(
	paymentRepo repository.PaymentsRepository,
	paymentMethodRepo repository.PaymentMethodsRepository,
	bookingRepo repository.BookingsRepository,
	logger *zap.Logger,
) PaymentUsecase {
	return &paymentUsecase{
		paymentRepo:       paymentRepo,
		paymentMethodRepo: paymentMethodRepo,
		bookingRepo:       bookingRepo,
		logger:            logger,
	}
}

func (u *paymentUsecase) GetPaymentMethods() ([]dto.PaymentMethodResponse, error) {
	paymentMethods, err := u.paymentMethodRepo.GetAllPaymentMethods()
	if err != nil {
		u.logger.Error("failed to get payment methods", zap.Error(err))
		return nil, err
	}

	var responses []dto.PaymentMethodResponse
	for _, pm := range paymentMethods {
		logoURL := ""
		if pm.LogoUrl != nil {
			logoURL = *pm.LogoUrl
		}
		responses = append(responses, dto.PaymentMethodResponse{
			ID:       pm.Id,
			Name:     pm.Name,
			LogoURL:  logoURL,
			IsActive: pm.IsActive,
		})
	}

	return responses, nil
}

func (u *paymentUsecase) ProcessPayment(req dto.PaymentRequest) (*dto.PaymentResponse, error) {
	// Get booking
	booking, err := u.bookingRepo.GetBookingByID(req.BookingID)
	if err != nil {
		u.logger.Error("booking not found", zap.Int("booking_id", req.BookingID), zap.Error(err))
		return nil, errors.New("booking not found")
	}

	// Check if booking is still pending
	if booking.Status != "pending" {
		return nil, errors.New("booking is not in pending status")
	}

	// Verify payment method exists
	_, err = u.paymentMethodRepo.GetPaymentMethodByID(req.PaymentMethodID)
	if err != nil {
		return nil, errors.New("invalid payment method")
	}

	// Create payment record
	payment := &entity.PaymentsRepository{
		BookingId:       req.BookingID,
		PaymentMethodId: req.PaymentMethodID,
		Amount:          booking.TotalPrice,
		TransactionTime: time.Now(),
		Status:          "success",
	}

	err = u.paymentRepo.CreatePayment(payment)
	if err != nil {
		u.logger.Error("failed to create payment", zap.Error(err))
		return nil, errors.New("failed to process payment")
	}

	// Update booking status
	err = u.bookingRepo.UpdateBookingStatus(req.BookingID, "confirmed")
	if err != nil {
		u.logger.Error("failed to update booking status", zap.Error(err))
		return nil, errors.New("payment processed but failed to update booking")
	}

	return &dto.PaymentResponse{
		PaymentID:       payment.Id,
		BookingID:       booking.Id,
		Amount:          payment.Amount,
		PaymentMethod:   "Payment Method",
		Status:          payment.Status,
		TransactionTime: payment.TransactionTime.Format("2006-01-02 15:04:05"),
	}, nil
}
