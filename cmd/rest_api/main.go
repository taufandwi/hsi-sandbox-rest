package main

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
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
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
)

type Config struct {
	Server struct {
		Rest string `validate:"required" mapstructure:"rest"`
	} `mapstructure:"server"`
	DB struct {
		Driver   string `validate:"required" mapstructure:"driver"`
		Host     string `validate:"required" mapstructure:"host"`
		Port     int    `validate:"required" mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		DbName   string `validate:"required" mapstructure:"dbname"`
		SSLMode  string `mapstructure:"sslmode"`
		PoolSize uint64 `mapstructure:"pool_size"`
	} `mapstructure:"database"`
	JWT struct {
		JwtKey string `mapstructure:"jwt_key"`
	} `mapstructure:"jwt"`
	Logger struct {
		Path         string `validate:"required" mapstructure:"path"`
		MaxAge       int    `validate:"required,gte=1,lte=365" mapstructure:"max_age"`
		RotationTime int    `validate:"required,gte=1,lte=365" mapstructure:"rotation_time"`
	} `mapstructure:"logger"`
}

type jwtCustomClaims struct {
	Username string `json:"name"`
	UserID   uint64 `json:"user_id"`
	jwt.RegisteredClaims
}

func main() {
	//	------------ config ------------
	viper.SetConfigName("config-dev")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic("Error reading config file: " + err.Error())
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		panic("Unable to decode into struct: " + err.Error())
	}

	// --------- db ---------
	var db *gorm.DB
	// close database connection on exit
	defer func() {
		if db != nil {
			sqlDB, _ := db.DB()
			sqlDB.Close()
		}
	}()

	db, err = gorm.Open(
		postgres.Open(
			fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=public TimeZone=Asia/Jakarta",
				config.DB.Host,
				config.DB.Username,
				config.DB.Password,
				config.DB.DbName,
				config.DB.Port,
				config.DB.SSLMode,
			),
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
	userHandler := userHdl.NewHandler(userService, config.JWT.JwtKey)
	employeeHandler := employeeHdl.NewHandler(employeeService)

	// -------- api --------
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.BodyLimit("50M"))
	e.Use(middleware.Recover())

	// -------- config echo middleware for JWT -----------
	configJWT := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey:    []byte(config.JWT.JwtKey),
		SigningMethod: "HS256",
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": "please check your token"})
		},
	}

	// echo group
	eg := e.Group("/api/v1")

	// ----- register routes -----
	health_check.RegisterPath(eg)
	userHandler.RegisterPath(eg)
	employeeHandler.RegisterPath(eg, echojwt.WithConfig(configJWT))

	// Start server
	go func() {
		if err := e.Start(config.Server.Rest); err != nil {
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
