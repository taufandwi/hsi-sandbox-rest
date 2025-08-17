package repository

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
)

type User interface {
	CreateUser(ctx context.Context, u model.User) (err error)
	GetAllUsers(ctx context.Context) (users []model.User, err error)
	UpdateUser(ctx context.Context, u model.User) (err error)
	GetUserByUserName(ctx context.Context, username string) (user model.User, err error)
}
