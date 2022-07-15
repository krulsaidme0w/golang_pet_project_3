package user

import (
	"github.com/krulsaidme0w/golang_pet_project_3/internal/models"
)

type UseCase interface {
	Save(userRequest *models.UserRequest) error
	Get(id string) (*models.User, error)
	Update(user models.User, updatedUser models.UserRequest) error
	Delete(id string) error
}
