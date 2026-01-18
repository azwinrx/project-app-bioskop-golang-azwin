package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)

type PaymentsRepository interface {
	CreatePayment(payment *entity.PaymentsRepository) error
	GetPaymentByBookingID(bookingID int) (*entity.PaymentsRepository, error)
}

type paymentsRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewPaymentsRepository(db database.PgxIface, log *zap.Logger) PaymentsRepository {
	return &paymentsRepository{db: db, logger: log}
}

func (r *paymentsRepository) CreatePayment(payment *entity.PaymentsRepository) error {
	query := `
		INSERT INTO payments (booking_id, payment_method_id, amount, transaction_time, status)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	
	err := r.db.QueryRow(
		context.Background(), 
		query, 
		payment.BookingId, 
		payment.PaymentMethodId, 
		payment.Amount,
		payment.TransactionTime,
		payment.Status,
	).Scan(&payment.Id)
	
	if err != nil {
		r.logger.Error("failed to create payment", 
			zap.Int("booking_id", payment.BookingId),
			zap.Error(err),
		)
		return err
	}
	
	r.logger.Info("payment created successfully", 
		zap.Int("payment_id", payment.Id),
		zap.Int("booking_id", payment.BookingId),
	)
	
	return nil
}

func (r *paymentsRepository) GetPaymentByBookingID(bookingID int) (*entity.PaymentsRepository, error) {
	query := `
		SELECT id, booking_id, payment_method_id, amount, transaction_time, status
		FROM payments
		WHERE booking_id = $1
	`
	
	payment := &entity.PaymentsRepository{}
	err := r.db.QueryRow(context.Background(), query, bookingID).Scan(
		&payment.Id,
		&payment.BookingId,
		&payment.PaymentMethodId,
		&payment.Amount,
		&payment.TransactionTime,
		&payment.Status,
	)
	
	if err != nil {
		r.logger.Error("failed to get payment by booking ID", 
			zap.Int("booking_id", bookingID),
			zap.Error(err),
		)
		return nil, err
	}
	
	return payment, nil
}
