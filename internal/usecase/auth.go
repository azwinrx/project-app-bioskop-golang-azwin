package usecase

import (
	"context"
	"errors"
	"fmt"
	"project-app-bioskop-golang-azwin/internal/data/entity"
	"project-app-bioskop-golang-azwin/internal/data/repository"
	"project-app-bioskop-golang-azwin/internal/dto"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(req dto.RegisterRequest) (*dto.AuthResponse, error)
	Login(req dto.LoginRequest) (*dto.LoginResponse, error)
	Logout(sessionID uuid.UUID) error
}

type authUsecase struct {
	userRepo repository.UsersRepository
	logger   *zap.Logger
}

func NewAuthUsecase(userRepo repository.UsersRepository, logger *zap.Logger) AuthUsecase {
	return &authUsecase{
		userRepo: userRepo,
		logger:   logger,
	}
}

func (u *authUsecase) Register(req dto.RegisterRequest) (*dto.AuthResponse, error) {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		u.logger.Error("failed to hash password", zap.Error(err))
		return nil, errors.New("failed to process password")
	}

	// Create user entity
	user := &entity.UsersRepository{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// Save to database
	err = u.userRepo.RegisterUsers(user)
	if err != nil {
		u.logger.Error("failed to register user", zap.Error(err))
		return nil, errors.New("failed to register user, email might already exist")
	}

	return &dto.AuthResponse{
		ID:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (u *authUsecase) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	ctx := context.Background()
	
	// Get user by email
	user, err := u.userRepo.LoginUsers(req.Email)
	if err != nil {
		u.logger.Error("user not found", zap.String("email", req.Email), zap.Error(err))
		return nil, errors.New("invalid email or password")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		u.logger.Warn("invalid password attempt", zap.String("email", req.Email))
		return nil, errors.New("invalid email or password")
	}

	// Create session with UUID
	sessionID := uuid.New()
	expiresAt := time.Now().Add(24 * time.Hour)
	
	session := &entity.Session{
		ID:        sessionID,
		UserID:    user.Id,
		ExpiresAt: expiresAt,
		CreatedAt: time.Now(),
	}

	err = u.userRepo.CreateSession(ctx, session)
	if err != nil {
		u.logger.Error("failed to create session", zap.Error(err))
		return nil, errors.New("failed to create session")
	}

	return &dto.LoginResponse{
		Token:     sessionID,
		ExpiresAt: expiresAt,
		User: dto.AuthResponse{
			ID:       user.Id,
			Username: user.Username,
			Email:    user.Email,
		},
	}, nil
}

func (u *authUsecase) Logout(sessionID uuid.UUID) error {
	ctx := context.Background()
	
	err := u.userRepo.RevokeSession(ctx, sessionID)
	if err != nil {
		u.logger.Error("failed to revoke session", zap.String("session_id", sessionID.String()), zap.Error(err))
		return fmt.Errorf("failed to logout")
	}
	return nil
}
