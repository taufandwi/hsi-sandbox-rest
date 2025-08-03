package request

type UserInput struct {
	Fullname string `json:"fullname" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}
