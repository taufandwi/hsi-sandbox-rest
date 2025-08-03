package user

import (
	"github.com/labstack/echo/v4"
	"github.com/taufandwi/hsi-sandbox-rest/handler/user/request"
	"github.com/taufandwi/hsi-sandbox-rest/handler/user/response"
	"github.com/taufandwi/hsi-sandbox-rest/service/user"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
)

type Handler struct {
	UserService user.Service
	// employeeService employee.Service // Example of another service that could be injected
}

func NewHandler(userService user.Service) *Handler {
	return &Handler{
		UserService: userService,
	}
}

// RegisterPath registers the user-related routes
func (h *Handler) RegisterPath(e *echo.Group) {
	e.POST("/user/create", h.CreateUser)
	e.GET("/user/get-all", h.GetAllUsers)
}

// CreateUser creates a new user
func (h *Handler) CreateUser(c echo.Context) error {
	var userInput request.UserInput
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if len(userInput.Fullname) > 100 {
		return c.JSON(400, map[string]string{"error": "Fullname exceeds maximum length of 100 characters"})
	}

	newUser := model.User{
		Fullname: userInput.Fullname,
		Email:    userInput.Email,
	}

	if err := h.UserService.CreateUser(newUser); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, newUser)
}

// GetAllUsers retrieves all users
func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve users"})
	}

	if len(users) == 0 {
		return c.JSON(204, nil) // No content
	}

	var userList []response.User
	for _, item := range users {
		userList = append(userList, response.User{
			ID:       item.ID,
			Fullname: item.Fullname,
			Email:    item.Email,
		})
	}

	return c.JSON(200, userList)
}
