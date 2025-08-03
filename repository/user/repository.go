package user

import "github.com/taufandwi/hsi-sandbox-rest/service/user/model"

type Repository struct {
	ModelUserList *[]model.User
}

func NewRepository(modelUserList *[]model.User) *Repository {
	return &Repository{
		ModelUserList: modelUserList,
	}
}

func (r *Repository) CreateUser(u model.User) (err error) {
	*r.ModelUserList = append(*r.ModelUserList, u)
	return nil
}

func (r *Repository) GetAllUsers() (users []model.User, err error) {
	if r.ModelUserList == nil {
		return nil, nil
	}
	return *r.ModelUserList, nil
}
