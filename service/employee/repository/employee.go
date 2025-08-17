package repository

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
)

type Employee interface {
	CreateEmployee(ctx context.Context, e model.Employee) (err error)
	GetAllEmployees(ctx context.Context) (employees []model.Employee, err error)
	UpdateEmployee(ctx context.Context, id int64, e model.Employee) (err error)
}
