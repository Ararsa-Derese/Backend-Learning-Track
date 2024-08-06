package router

import (
	"taskmanager/controllers"
	"taskmanager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	middleware.InitJWTSecret()
	router := gin.Default()

	// Public routes
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)

	// Apply middleware for subsequent routes
	router.Use(middleware.AuthMiddleware())

	// Protected routes
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTask)

	// admin only routes
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.POST("/tasks", controllers.Addtask)

	return router
}
