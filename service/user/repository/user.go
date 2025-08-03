package repository

import "github.com/taufandwi/hsi-sandbox-rest/service/user/model"

type User interface {
	CreateUser(u model.User) (err error)
	GetAllUsers() (users []model.User, err error)
	UpdateUser(u model.User) (err error)
}
