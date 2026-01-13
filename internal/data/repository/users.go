package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)



type UsersRepository interface {

}

type usersRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewUsersRepository(db database.PgxIface, log *zap.Logger) UsersRepository {
	return &usersRepository{db: db, Logger: log}
}

func (r *usersRepository) RegisterUsers(data *entity.UsersRepository) error {
	query := `
		INSERT INTO users (username, email, password)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := r.db.QueryRow(context.Background(), query, data.Username, data.Email, data.Password).Scan(&data.Id)
	if err != nil {
		r.Logger.Error("failed to create user",
			zap.String("username", data.Username),
			zap.String("email", data.Email),
			zap.Error(err),
		)
		return err
	}
	r.Logger.Info("user created successfully",
		zap.Int("user_id", data.Id),
		zap.String("username", data.Username),
		zap.String("email", data.Email),
	)
	return nil
}

func (r *usersRepository) LoginUsers(email string) (*entity.UsersRepository, error)