package health_check

import "github.com/labstack/echo/v4"

func RegisterPath(e *echo.Group) {
	e.GET("/ping", healthCheck)
}

// healthCheck is a simple handler to check if the server is running
func healthCheck(c echo.Context) error {
	var respons struct {
		Message string `json:"message"`
	}

	respons.Message = "pong"

	return c.JSON(200, respons)
}
