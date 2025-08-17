package employee

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/repository"
)

type Service interface {
	CreateEmployee(ctx context.Context, e model.Employee) (err error)
	GetAllEmployees(ctx context.Context) (employees []model.Employee, err error)
	UpdateEmployee(ctx context.Context, id int64, e model.Employee) (err error)
}

type service struct {
	employeeRepo repository.Employee
}

func NewService(employeeRepo repository.Employee) Service {
	return &service{employeeRepo}
}

func (s *service) CreateEmployee(ctx context.Context, e model.Employee) (err error) {
	return s.employeeRepo.CreateEmployee(ctx, e)
}

func (s *service) GetAllEmployees(ctx context.Context) (employees []model.Employee, err error) {
	return s.employeeRepo.GetAllEmployees(ctx)
}

func (s *service) UpdateEmployee(ctx context.Context, id int64, e model.Employee) (err error) {
	return s.employeeRepo.UpdateEmployee(ctx, id, e)
}
