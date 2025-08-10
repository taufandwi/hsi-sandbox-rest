package response

import "github.com/taufandwi/hsi-sandbox-rest/service/user/model"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func NewUserResponse(u model.User) User {
	return User{
		ID:       u.ID,
		Username: u.Username,
	}
}
