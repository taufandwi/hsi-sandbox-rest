package employee

import (
	"context"
	"github.com/taufandwi/hsi-sandbox-rest/repository/employee/entity"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) CreateEmployee(ctx context.Context, e model.Employee) (err error) {
	tx := r.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte("test123"), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	// create in table users
	userEnt := entity.User{
		Username:     e.Email,
		PasswordHash: string(hashPass),
	}

	if err = tx.Create(&userEnt).Error; err != nil {
		tx.Rollback()
		return
	}

	employeeEnt := entity.NewEmployeeEntity(e)
	employeeEnt.UserID = userEnt.ID

	if err = tx.Create(&employeeEnt).Error; err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
	}

	return
}

func (r *Repository) GetAllEmployees(ctx context.Context) (employees []model.Employee, err error) {
	var employeeEnts []entity.Employee

	if err = r.db.WithContext(ctx).Find(&employeeEnts).Error; err != nil {
		return
	}

	for _, item := range employeeEnts {
		employees = append(employees, item.ToModel())
	}

	return
}

func (r *Repository) UpdateEmployee(ctx context.Context, id int64, e model.Employee) (err error) {
	employeeEnt := entity.NewEmployeeEntity(e)

	if err = r.db.WithContext(ctx).Omit("id", "user_id", "email", "phone_number", "hire_date", "department").Where("id = ?", id).Updates(&employeeEnt).Error; err != nil {
		return
	}

	return
}
