package usecase

import (
	"project-app-bioskop-golang-azwin/internal/data/repository"
)

type UserUsecase struct {
	repo *repository.UsersRepository
	
}