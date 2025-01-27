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

	// Task setup
	taskRepository := repositories.NewTaskRepository(db)
	taskService := services.NewTaskService(taskRepository)
	taskHandler := handlers.NewTaskHandler(taskService)

	// Comment setup
	commentRepository := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository)
	commentHandler := handlers.NewCommentHandler(commentService)

	// Enum setup
	enumService := services.NewEnumService()
	enumHandler := handlers.NewEnumHandler(enumService)

	// Role setup
	roleRepository := repositories.NewRoleRepository(db)
	roleService := services.NewRoleService(roleRepository)
	roleHandler := handlers.NewRoleHandler(roleService)

	api := e.Group("/api")

	// User routes
	api.POST("/users", userHandler.Store)
	api.GET("/users", userHandler.Index)
	api.GET("/users/:id", userHandler.Show)
	api.PUT("/users/:id", userHandler.Edit)
	api.DELETE("/users/:id", userHandler.Destroy)

	// Task routes
	api.POST("/tasks", taskHandler.Store)
	api.GET("/tasks", taskHandler.Index)
	api.GET("/tasks/:id", taskHandler.Show)
	api.PUT("/tasks/:id", taskHandler.Edit)
	api.DELETE("/tasks/:id", taskHandler.Destroy)
	api.PUT("/tasks/:id/status", taskHandler.EditStatus)

	// Comment routes
	api.POST("/comments", commentHandler.Store)
	api.GET("/comments", commentHandler.Index)
	api.GET("/comments/:id", commentHandler.Show)
	api.PUT("/comments/:id", commentHandler.Edit)
	api.DELETE("/comments/:id", commentHandler.Destroy)

	// Enum routes
	api.GET("/enums/task_statuses", enumHandler.GetTaskStatuses)

	// Role routes
	api.POST("/roles", roleHandler.Create)
	api.GET("/roles", roleHandler.Index)
	api.GET("/roles/:id", roleHandler.Show)
	api.PUT("/roles/:id", roleHandler.Edit)
	api.DELETE("/roles/:id", roleHandler.Destroy)
}
