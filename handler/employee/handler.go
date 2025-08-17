package employee

import (
	"github.com/labstack/echo/v4"
	"github.com/taufandwi/hsi-sandbox-rest/handler/employee/request"
	"github.com/taufandwi/hsi-sandbox-rest/handler/employee/response"
	"github.com/taufandwi/hsi-sandbox-rest/service/employee"
	"strconv"
)

type Handler struct {
	EmployeeService employee.Service
}

func NewHandler(employeeService employee.Service) *Handler {
	return &Handler{
		employeeService,
	}
}

// RegisterPath registers the employee-related routes
func (h *Handler) RegisterPath(e *echo.Group, middlewareJWt ...echo.MiddlewareFunc) {
	e.POST("/employee/create", h.createEmployee)
	e.GET("/employee/get-all", h.getAllEmployees)
	e.PUT("/employee/update", h.updateEmployee)
}

// createEmployee creates a new employee
func (h *Handler) createEmployee(c echo.Context) error {
	var employeeReq request.Employee

	if err := c.Bind(&employeeReq); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	if err := h.EmployeeService.CreateEmployee(c.Request().Context(), employeeReq.ToModel()); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create employee"})
	}

	return c.JSON(201, map[string]string{"message": "Employee created successfully"})
}

// getAllEmployees retrieves all employees
func (h *Handler) getAllEmployees(c echo.Context) error {
	employees, err := h.EmployeeService.GetAllEmployees(c.Request().Context())
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to retrieve employees"})
	}

	if len(employees) == 0 {
		return c.JSON(204, nil) // No content
	}

	// Convert the employee models to response format
	var employeeResponses []response.Employee
	for _, item := range employees {
		employeeResponses = append(employeeResponses, response.NewEmployeeResponse(item))
	}

	return c.JSON(200, employeeResponses)
}

// updateEmployee updates an existing employee
func (h *Handler) updateEmployee(c echo.Context) error {
	var employeeReq request.Employee

	if err := c.Bind(&employeeReq); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid input"})
	}

	idStr := c.QueryParam("id")
	if idStr == "" {
		return c.JSON(400, map[string]string{"error": "Employee ID is required"})
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid employee ID"})
	}

	if err := h.EmployeeService.UpdateEmployee(c.Request().Context(), id, employeeReq.ToModel()); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to update employee"})
	}

	return c.JSON(200, map[string]string{"message": "Employee updated successfully"})
}
