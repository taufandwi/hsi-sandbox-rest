package user

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/taufandwi/hsi-sandbox-rest/handler/user/request"
	"github.com/taufandwi/hsi-sandbox-rest/handler/user/response"
	"github.com/taufandwi/hsi-sandbox-rest/service/user"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
	"golang.org/x/crypto/bcrypt"
	"time"
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
	e.POST("/user/login", h.LoginAndGenerateJWTToken)
}

// CreateUser creates a new user
func (h *Handler) CreateUser(c echo.Context) error {
	var userInput request.UserInput
	if err := c.Bind(&userInput); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if len(userInput.Username) > 100 {
		return c.JSON(400, map[string]string{"error": "Fullname exceeds maximum length of 100 characters"})
	}

	newUser := model.User{
		Username: userInput.Username,
		Password: userInput.Password,
	}

	if err := h.UserService.CreateUser(newUser); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, map[string]string{"message": "User created successfully"})
}

// GetAllUsers retrieves all users
func (h *Handler) GetAllUsers(c echo.Context) error {
	users, err := h.UserService.GetAllUser()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve users"})
	}

	if len(users) == 0 {
		return c.JSON(204, nil) // No content
	}

	var userList []response.User
	for _, item := range users {
		userList = append(userList, response.NewUserResponse(item))
	}

	return c.JSON(200, userList)
}

func (h *Handler) LoginAndGenerateJWTToken(c echo.Context) error {
	var loginInfo request.LoginInfo
	if err := c.Bind(&loginInfo); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	userMdl, err := h.UserService.GetUserByUserName(c.Request().Context(), loginInfo.Username)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	err = bcrypt.CompareHashAndPassword([]byte(userMdl.Password), []byte(loginInfo.Password))
	if err != nil {
		return c.JSON(401, map[string]string{"error": "Invalid credentials"})
	}

	// Set custom claims
	claims := &jwtCustomClaims{
		userMdl.Username,
		uint64(userMdl.ID),
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}

	// Create token with claims
	tokenTemp := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	token, err := tokenTemp.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(200, map[string]string{"token": token})

}

type jwtCustomClaims struct {
	Username string `json:"name"`
	UserID   uint64 `json:"user_id"`
	jwt.RegisteredClaims
}
