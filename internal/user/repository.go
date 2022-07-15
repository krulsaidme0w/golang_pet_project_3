package user

import (
	"github.com/krulsaidme0w/golang_pet_project_3/internal/models"
)

type Repository interface {
	Save(user *models.User) error
	Get(id string) (*models.User, error)
	Update(user *models.User) error
	Delete(id string) error
}
