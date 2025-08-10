package repository

import "github.com/taufandwi/hsi-sandbox-rest/service/employee/model"

type Employee interface {
	CreateEmployee(e model.Employee) (err error)
	GetAllEmployees() (employees []model.Employee, err error)
	UpdateEmployee(id int64, e model.Employee) (err error)
}
