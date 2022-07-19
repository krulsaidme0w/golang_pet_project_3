package server

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/krulsaidme0w/golang_pet_project_3/pkg/security"
	userservice "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service"
	proto "github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/grpc/proto"
	"github.com/krulsaidme0w/golang_pet_project_3/pkg/user-service/models"
)

type userService struct {
	userRepo userservice.UserRepository
	proto.UnimplementedUserUseCaseServer
}

func NewUserService(userRepo userservice.UserRepository) proto.UserUseCaseServer {
	return &userService{
		userRepo: userRepo,
	}
}

func (u *userService) Save(ctx context.Context, user *proto.User) (*emptypb.Empty, error) {
	err := u.userRepo.Save(ctx, getUserForRepo(user))

	return &emptypb.Empty{}, err
}

func (u *userService) Get(ctx context.Context, id *proto.UserID) (*proto.User, error) {
	user, err := u.userRepo.Get(ctx, id.GetId())
	if err != nil {
		return nil, err
	}

	return getProtoUser(user), err
}

func (u *userService) Update(ctx context.Context, user *proto.User) (*emptypb.Empty, error) {
	userForRepo := getUserForRepo(user)
	userForRepo.Password = security.Hash(userForRepo.Password)

	err := u.userRepo.Update(ctx, userForRepo)

	return &emptypb.Empty{}, err
}

func (u *userService) Delete(ctx context.Context, id *proto.UserID) (*emptypb.Empty, error) {
	err := u.userRepo.Delete(ctx, id.GetId())

	return &emptypb.Empty{}, err
}

func getUserForRepo(userProto *proto.User) *models.User {
	return &models.User{
		ID:       userProto.GetId(),
		Username: userProto.GetUsername(),
		Email:    userProto.GetEmail(),
		Password: userProto.GetPassword(),
	}
}

func getProtoUser(user *models.User) *proto.User {
	return &proto.User{
		Id:       uint64(user.ID),
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}
}
