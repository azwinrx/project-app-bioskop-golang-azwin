package repository

import (
	"context"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/pkg/database"

	"github.com/google/uuid"
	"go.uber.org/zap"
)



type UsersRepository interface {
	RegisterUsers(data *entity.UsersRepository) error
	LoginUsers(email string) (*entity.UsersRepository, error)
	LogoutUsers(userId int) error
	
	// Session methods
	CreateSession(ctx context.Context, session *entity.Session) error
	ValidateSession(ctx context.Context, sessionID uuid.UUID) (*entity.Session, error)
	GetSessionByUserId(ctx context.Context, userId int) (*entity.Session, error)
	RevokeSession(ctx context.Context, sessionID uuid.UUID) error
	RevokeAllUserSessions(ctx context.Context, userId int) error
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

// CreateSession creates a new session for a user
func (r *usersRepository) CreateSession(ctx context.Context, session *entity.Session) error {
	query := `
		INSERT INTO sessions (id, user_id, expires_at, created_at)
		VALUES ($1, $2, $3, $4)
	`
	
	_, err := r.db.Exec(ctx, query, 
		session.ID,
		session.UserID,
		session.ExpiresAt,
		session.CreatedAt,
	)
	
	if err != nil {
		r.Logger.Error("failed to create session",
			zap.String("session_id", session.ID.String()),
			zap.Int("user_id", session.UserID),
			zap.Error(err),
		)
		return err
	}
	
	r.Logger.Info("session created successfully",
		zap.String("session_id", session.ID.String()),
		zap.Int("user_id", session.UserID),
	)
	
	return nil
}

// ValidateSession validates a session token and returns session details if valid
func (r *usersRepository) ValidateSession(ctx context.Context, sessionID uuid.UUID) (*entity.Session, error) {
	query := `
		SELECT id, user_id, expires_at, revoked_at, created_at
		FROM sessions
		WHERE id = $1
		  AND revoked_at IS NULL
		  AND expires_at > NOW()
	`
	
	session := &entity.Session{}
	err := r.db.QueryRow(ctx, query, sessionID).Scan(
		&session.ID,
		&session.UserID,
		&session.ExpiresAt,
		&session.RevokedAt,
		&session.CreatedAt,
	)
	
	if err != nil {
		r.Logger.Error("failed to validate session",
			zap.String("session_id", sessionID.String()),
			zap.Error(err),
		)
		return nil, err
	}
	
	r.Logger.Info("session validated successfully",
		zap.String("session_id", session.ID.String()),
		zap.Int("user_id", session.UserID),
	)

	return session, nil
}

// GetSessionByUserId retrieves the latest active session for a user
func (r *usersRepository) GetSessionByUserId(ctx context.Context, userId int) (*entity.Session, error) {
	query := `
		SELECT id, user_id, expires_at, revoked_at, created_at
		FROM sessions
		WHERE user_id = $1
		  AND revoked_at IS NULL
		  AND expires_at > NOW()
		ORDER BY created_at DESC
		LIMIT 1
	`
	
	session := &entity.Session{}
	err := r.db.QueryRow(ctx, query, userId).Scan(
		&session.ID,
		&session.UserID,
		&session.ExpiresAt,
		&session.RevokedAt,
		&session.CreatedAt,
	)
	
	if err != nil {
		r.Logger.Error("failed to get session by user id",
			zap.Int("user_id", userId),
			zap.Error(err),
		)
		return nil, err
	}
	
	return session, nil
}

// RevokeSession revokes a specific session
func (r *usersRepository) RevokeSession(ctx context.Context, sessionID uuid.UUID) error {
	query := `
		UPDATE sessions
		SET revoked_at = NOW()
		WHERE id = $1
		  AND revoked_at IS NULL
	`
	
	_, err := r.db.Exec(ctx, query, sessionID)
	if err != nil {
		r.Logger.Error("failed to revoke session",
			zap.String("session_id", sessionID.String()),
			zap.Error(err),
		)
		return err
	}
	
	r.Logger.Info("session revoked successfully",
		zap.String("session_id", sessionID.String()),
	)
	
	return nil
}

// RevokeAllUserSessions revokes all sessions for a specific user
func (r *usersRepository) RevokeAllUserSessions(ctx context.Context, userId int) error {
	query := `
		UPDATE sessions
		SET revoked_at = NOW()
		WHERE user_id = $1
		  AND revoked_at IS NULL
		  AND expires_at > NOW()
	`
	
	_, err := r.db.Exec(ctx, query, userId)
	if err != nil {
		r.Logger.Error("failed to revoke all user sessions",
			zap.Int("user_id", userId),
			zap.Error(err),
		)
		return err
	}
	
	r.Logger.Info("all user sessions revoked successfully",
		zap.Int("user_id", userId),
	)
	
	return nil
}