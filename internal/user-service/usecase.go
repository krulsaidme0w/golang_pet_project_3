package userservice

import (
	"context"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type UserUseCase interface {
	Save(ctx context.Context, userRequest *models.UserRequest) error
	Get(ctx context.Context, id string) (*models.User, error)
	Update(ctx context.Context, user *models.User, updatedUser *models.UserRequest) error
	Delete(ctx context.Context, id string) error
}
