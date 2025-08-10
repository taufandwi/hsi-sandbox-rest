package entity

import "github.com/taufandwi/hsi-sandbox-rest/service/user/model"

type User struct {
	ID           int64 `gorm:"primaryKey"`
	Username     string
	PasswordHash string
}

func (u User) ToModel() model.User {
	return model.User{
		ID:       u.ID,
		Username: u.Username,
	}
}
