package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	employeeHdl "github.com/taufandwi/hsi-sandbox-rest/handler/employee"
	"github.com/taufandwi/hsi-sandbox-rest/handler/health_check"
	userHdl "github.com/taufandwi/hsi-sandbox-rest/handler/user"
	"github.com/taufandwi/hsi-sandbox-rest/repository/employee"
	"github.com/taufandwi/hsi-sandbox-rest/repository/user"
	employeeSrv "github.com/taufandwi/hsi-sandbox-rest/service/employee"
	userSrv "github.com/taufandwi/hsi-sandbox-rest/service/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"os/signal"
	"time"
)

func main() {

	// --------- db ---------
	var db *gorm.DB
	// close database connection on exit
	defer func() {
		if db != nil {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}
	}()

	db, err := gorm.Open(
		postgres.Open(
			"host=localhost user=postgres password=postgres dbname=hsi_sandbox port=5433 sslmode=disable TimeZone=Asia/Jakarta",
		),
		&gorm.Config{
			// debug mode
			Logger: logger.Default.LogMode(logger.Info),
			//Logger: logger.Default.LogMode(logger.Silent),
		},
	)
	if err != nil {
		panic(err)
	}

	// -------- init repo --------
	userRepository := user.NewRepository(db)
	employeeRepository := employee.NewRepository(db)

	// -------- init service --------
	userService := userSrv.NewService(userRepository)
	employeeService := employeeSrv.NewService(employeeRepository)

	// -------- init handler --------
	userHandler := userHdl.NewHandler(userService)
	employeeHandler := employeeHdl.NewHandler(employeeService)

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
	employeeHandler.RegisterPath(eg)

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
