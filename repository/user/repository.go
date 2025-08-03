package user

import "github.com/taufandwi/hsi-sandbox-rest/service/user/model"

// implementation of the User repository interface
type Repository struct {
	// mock for database
	ModelUserList *[]model.User

	// orm
	// DB *gorm.DB, DB *sqlx.DB

	// other connections
	// Cache *redis.Client
	// Mongo *mongo.Client
}

func NewRepository(modelUserList *[]model.User) *Repository {
	return &Repository{
		ModelUserList: modelUserList,
	}
}

func (r *Repository) CreateUser(u model.User) (err error) {
	u.ID = int64(len(*r.ModelUserList) + 1) // Simple ID generation logic, can be replaced with a better one
	*r.ModelUserList = append(*r.ModelUserList, u)
	return nil
}

func (r *Repository) GetAllUsers() (users []model.User, err error) {
	// gorm.find(&users)
	// sqlx.Select(&users, "SELECT * FROM users")
	if r.ModelUserList == nil {
		return nil, nil
	}

	return *r.ModelUserList, nil
}

func (r *Repository) UpdateUser(u model.User) (err error) {
	for i, user := range *r.ModelUserList {
		if user.ID == u.ID {
			(*r.ModelUserList)[i] = u
			return nil
		}
	}
	return nil // or return an error if user not found
}
