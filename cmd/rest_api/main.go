package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taufandwi/hsi-sandbox-rest/handler/health_check"
	userRepo "github.com/taufandwi/hsi-sandbox-rest/handler/user"
	"github.com/taufandwi/hsi-sandbox-rest/repository/user"
	modelEmployee "github.com/taufandwi/hsi-sandbox-rest/service/employee/model"
	userSrv "github.com/taufandwi/hsi-sandbox-rest/service/user"
	"github.com/taufandwi/hsi-sandbox-rest/service/user/model"
	"os"
	"os/signal"
	"time"
)

func main() {

	// init awal, orm, cache, etc.
	// Create a mockup data store for users
	modelUserList := make([]model.User, 0)
	_ = make([]modelEmployee.Employee, 0)

	// -------- init repo --------
	userRepository := user.NewRepository(&modelUserList)

	// -------- init service --------
	userService := userSrv.NewService(userRepository)

	// -------- init handler --------
	userHandler := userRepo.NewHandler(userService)

	// -------- api --------
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("50M"))
	e.Use(middleware.Recover())

	// echo group
	eg := e.Group("/api/v1")

	// ----- register routes -----
	health_check.RegisterPath(eg)
	userHandler.RegisterPath(eg)

	// Start server
	go func() {
		if err := e.Start(":55667"); err != nil {
			e.Logger.Fatal("Shutting down the server")
		}
	}()

	fmt.Println("Sandbox HSI service is started")

	// Wait for shutdown signal for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	fmt.Println("Shutting down the server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal("Error shutting down the server:", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}
}
