package usecase

import (
	"context"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/security"
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type userUseCase struct {
	repository userservice.UserRepository
}

func NewUserUseCase(repository userservice.UserRepository) *userUseCase {
	return &userUseCase{
		repository: repository,
	}
}

func (u *userUseCase) Save(ctx context.Context, userRequest *models.UserRequest) error {
	user, err := userRequest.ToUser()
	if err != nil {
		return err
	}

	return u.repository.Save(ctx, user)
}

func (u *userUseCase) Get(ctx context.Context, id uint64) (*models.User, error) {
	return u.repository.Get(ctx, id)
}

func (u *userUseCase) Update(ctx context.Context, user *models.User) error {
	user.Password = security.Hash(user.Password)

	return u.repository.Update(ctx, user)
}

func (u *userUseCase) Delete(ctx context.Context, id uint64) error {
	return u.repository.Delete(ctx, id)
}
