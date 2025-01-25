package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/qsmsoft/todo/config"
	"github.com/qsmsoft/todo/internal/database"
	"github.com/qsmsoft/todo/internal/handlers"
	"github.com/qsmsoft/todo/internal/repositories"
	"github.com/qsmsoft/todo/internal/services"
)

func RegisterRoutes(e *echo.Echo) {

	db := database.NewDatabase(*config.LoadConfig())
	// User setup
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)

	api := e.Group("/api")

	// User routes
	api.POST("/users", userHandler.Store)
	api.GET("/users", userHandler.Index)
	api.GET("/users/:id", userHandler.Show)
	api.PUT("/users/:id", userHandler.Edit)
	api.DELETE("/users/:id", userHandler.Destroy)
}
