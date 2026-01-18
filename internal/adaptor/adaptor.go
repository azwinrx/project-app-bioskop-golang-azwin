package adaptor

import (
	"project-app-bioskop-golang-azwin/internal/usecase"

	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type Adaptor struct {
	*AuthAdaptor
	*CinemaAdaptor
	*BookingAdaptor
	*PaymentAdaptor
}

func NewAdaptor(
	authUsecase usecase.AuthUsecase,
	cinemaUsecase usecase.CinemaUsecase,
	bookingUsecase usecase.BookingUsecase,
	paymentUsecase usecase.PaymentUsecase,
	validator *validator.Validate,
	logger *zap.Logger,
) *Adaptor {
	return &Adaptor{
		AuthAdaptor:    NewAuthAdaptor(authUsecase, validator, logger),
		CinemaAdaptor:  NewCinemaAdaptor(cinemaUsecase, logger),
		BookingAdaptor: NewBookingAdaptor(bookingUsecase, validator, logger),
		PaymentAdaptor: NewPaymentAdaptor(paymentUsecase, validator, logger),
	}
}
