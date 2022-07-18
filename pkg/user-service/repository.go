package userservice

import (
	"context"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type UserRepository interface {
	Save(ctx context.Context, user *models.User) error
	Get(ctx context.Context, id uint64) (*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint64) error
}
