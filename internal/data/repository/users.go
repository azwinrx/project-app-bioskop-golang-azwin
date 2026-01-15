package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"go.uber.org/zap"
)



type UsersRepository interface {
	RegisterUsers(data *entity.UsersRepository) error
	LoginUsers(email string) (*entity.UsersRepository, error)
	LogoutUsers(userId int) error
}

type usersRepository struct {
	db database.PgxIface
	Logger *zap.Logger
}

func NewUsersRepository(db database.PgxIface, log *zap.Logger) UsersRepository {
	return &usersRepository{db: db, Logger: log}
}

// Function to register a new user
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

// Function to login
func (r *usersRepository) LoginUsers(email string) (*entity.UsersRepository, error){
	query := `
		SELECT id, username, email, password, is_verified, created_at, updated_at, deleted_at
		FROM users
		WHERE email = $1
	`
	
	user := &entity.UsersRepository{}
	err := r.db.QueryRow(context.Background(), query, email).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.IsVerified,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.DeletedAt,
	)
	
	if err != nil {
		r.Logger.Error("failed to find user for login",
			zap.String("email", email),
			zap.Error(err),
		)
		return nil, err
	}
	
	r.Logger.Info("user found for login",
		zap.Int("user_id", user.Id),
		zap.String("email", email),
	)
	
	return user, nil
}

// Function to logout
func (r *usersRepository) LogoutUsers(userId int) error {
	query := `
		UPDATE users
		SET updated_at = NOW()
		WHERE id = $1
	`
	
	_, err := r.db.Exec(context.Background(), query, userId)
	if err != nil {
		r.Logger.Error("failed to update user logout time",
			zap.Int("user_id", userId),
			zap.Error(err),
		)
		return err
	}
	
	r.Logger.Info("user logged out successfully",
		zap.Int("user_id", userId),
	)
	
	return nil
}