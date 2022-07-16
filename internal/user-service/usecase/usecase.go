package usecase

import (
	"github.com/krulsaidme0w/golang_pet_project_3/internal/user-service/repository"
)

type userUseCase struct {
	repository repository.UserRepository
}

func NewUserUseCase(repository repository.UserRepository) *userUseCase {
	return &userUseCase{
		repository: repository,
	}
}
