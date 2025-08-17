package user

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/repository"
)

type Service interface {
	CreateUser(ctx context.Context, u model.User) (err error)
	GetAllUser(ctx context.Context) (users []model.User, err error)
	UpdateUser(ctx context.Context, u model.User) (err error)
	GetUserByUserName(ctx context.Context, username string) (user model.User, err error)
}

type service struct {
	userRepo repository.User
}

func NewService(userRepo repository.User) Service {
	return &service{userRepo}
}

func (s *service) CreateUser(ctx context.Context, u model.User) (err error) {
	return s.userRepo.CreateUser(ctx, u)
}

func (s *service) GetAllUser(ctx context.Context) (users []model.User, err error) {
	return s.userRepo.GetAllUsers(ctx)
}

func (s *service) UpdateUser(ctx context.Context, u model.User) (err error) {
	return s.userRepo.UpdateUser(ctx, u)
}

func (s *service) GetUserByUserName(ctx context.Context, username string) (user model.User, err error) {
	return s.userRepo.GetUserByUserName(ctx, username)
}
