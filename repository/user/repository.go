package user

import (
	"github.com/taufandwi/hsi-sandbox-rest/repository/user/entity"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// implementation of the User repository interface
type Repository struct {
	// mock for database
	//ModelUserList *[]model.User

	// orm
	db *gorm.DB
	//DB *sqlx.DB

	// other connections
	// Cache *redis.Client
	// Mongo *mongo.Client
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db,
	}
}

func (r *Repository) CreateUser(u model.User) (err error) {
	var userEnt entity.User

	// hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}

	userEnt = entity.User{
		Username:     u.Username,
		PasswordHash: string(hashPassword),
	}

	// save to database
	if err = r.db.Create(&userEnt).Error; err != nil {
		return
	}

	return nil
}

func (r *Repository) GetAllUsers() (users []model.User, err error) {
	var userEnts []entity.User

	if err = r.db.Order("id desc").Find(&userEnts).Error; err != nil {
		return
	}

	for _, item := range userEnts {
		users = append(users, item.ToModel())
	}

	return
}

func (r *Repository) UpdateUser(u model.User) (err error) {
	//for i, user := range *r.ModelUserList {
	//	if user.ID == u.ID {
	//		(*r.ModelUserList)[i] = u
	//		return nil
	//	}
	//}
	return nil // or return an error if user not found
}
