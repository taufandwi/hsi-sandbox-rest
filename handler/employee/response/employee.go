package response

import (
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
)

type Employee struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	JobTitle    string `json:"job_title"`
	HireDate    string `json:"hire_date"`
	Department  string `json:"department"`
}

func NewEmployeeResponse(e model.Employee) Employee {
	return Employee{
		ID:          e.ID,
		UserID:      e.UserID,
		FirstName:   e.FirstName,
		LastName:    e.LastName,
		Email:       e.Email,
		PhoneNumber: e.PhoneNumber,
		JobTitle:    e.JobTitle,
		HireDate:    e.HireDate,
		Department:  e.Department,
	}
}
