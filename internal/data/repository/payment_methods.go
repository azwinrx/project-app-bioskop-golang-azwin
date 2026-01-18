package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)

type PaymentMethodsRepository interface {
	GetAllPaymentMethods() ([]*entity.PaymentMethodsRepository, error)
	GetPaymentMethodByID(paymentMethodID int) (*entity.PaymentMethodsRepository, error)
}

type paymentMethodsRepository struct {
	db     database.PgxIface
	logger *zap.Logger
}

func NewPaymentMethodsRepository(db database.PgxIface, log *zap.Logger) PaymentMethodsRepository {
	return &paymentMethodsRepository{db: db, logger: log}
}

func (r *paymentMethodsRepository) GetAllPaymentMethods() ([]*entity.PaymentMethodsRepository, error) {
	query := `
		SELECT id, name, logo_url, is_active, created_at
		FROM payment_methods
		WHERE is_active = true
		ORDER BY name
	`
	
	rows, err := r.db.Query(context.Background(), query)
	if err != nil {
		r.logger.Error("failed to get payment methods", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []*entity.PaymentMethodsRepository
	for rows.Next() {
		pm := &entity.PaymentMethodsRepository{}
		err := rows.Scan(
			&pm.Id,
			&pm.Name,
			&pm.LogoUrl,
			&pm.IsActive,
			&pm.CreatedAt,
		)
		if err != nil {
			r.logger.Error("failed to scan payment method", zap.Error(err))
			return nil, err
		}
		paymentMethods = append(paymentMethods, pm)
	}

	return paymentMethods, nil
}

func (r *paymentMethodsRepository) GetPaymentMethodByID(paymentMethodID int) (*entity.PaymentMethodsRepository, error) {
	query := `
		SELECT id, name, logo_url, is_active, created_at
		FROM payment_methods
		WHERE id = $1
	`
	
	pm := &entity.PaymentMethodsRepository{}
	err := r.db.QueryRow(context.Background(), query, paymentMethodID).Scan(
		&pm.Id,
		&pm.Name,
		&pm.LogoUrl,
		&pm.IsActive,
		&pm.CreatedAt,
	)
	
	if err != nil {
		r.logger.Error("failed to get payment method by ID", 
			zap.Int("payment_method_id", paymentMethodID),
			zap.Error(err),
		)
		return nil, err
	}
	
	return pm, nil
}
