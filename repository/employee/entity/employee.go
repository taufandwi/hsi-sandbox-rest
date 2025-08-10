package entity

import (
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
)

type Employee struct {
	ID          int64 `gorm:"primaryKey"`
	UserID      int64
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	JobTitle    string
	HireDate    string
	Department  string
}

func (e Employee) ToModel() model.Employee {
	return model.Employee{
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

func NewEmployeeEntity(emp model.Employee) Employee {
	return Employee{
		ID:          emp.ID,
		UserID:      emp.UserID,
		FirstName:   emp.FirstName,
		LastName:    emp.LastName,
		Email:       emp.Email,
		PhoneNumber: emp.PhoneNumber,
		JobTitle:    emp.JobTitle,
		HireDate:    emp.HireDate,
		Department:  emp.Department,
	}
}
