package user

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/repository"
)

type Service interface {
	CreateUser(u model.User) (err error)
	GetAllUser() (users []model.User, err error)
	UpdateUser(u model.User) (err error)
	GetUserByUserName(ctx context.Context, username string) (user model.User, err error)
}

type service struct {
	userRepo repository.User
}

func NewService(userRepo repository.User) Service {
	return &service{userRepo}
}

func (s *service) CreateUser(u model.User) (err error) {
	return s.userRepo.CreateUser(u)
}

func (s *service) GetAllUser() (users []model.User, err error) {
	return s.userRepo.GetAllUsers()
}

func (s *service) UpdateUser(u model.User) (err error) {
	return s.userRepo.UpdateUser(u)
}

func (s *service) GetUserByUserName(ctx context.Context, username string) (user model.User, err error) {
	return s.userRepo.GetUserByUserName(ctx, username)
}
