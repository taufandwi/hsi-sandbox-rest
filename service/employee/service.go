package employee

import (
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/repository"
)

type Service interface {
	CreateEmployee(e model.Employee) (err error)
	GetAllEmployees() (employees []model.Employee, err error)
	UpdateEmployee(id int64, e model.Employee) (err error)
}

type service struct {
	employeeRepo repository.Employee
}

func NewService(employeeRepo repository.Employee) Service {
	return &service{employeeRepo}
}

func (s *service) CreateEmployee(e model.Employee) (err error) {
	return s.employeeRepo.CreateEmployee(e)
}

func (s *service) GetAllEmployees() (employees []model.Employee, err error) {
	return s.employeeRepo.GetAllEmployees()
}

func (s *service) UpdateEmployee(id int64, e model.Employee) (err error) {
	return s.employeeRepo.UpdateEmployee(id, e)
}
