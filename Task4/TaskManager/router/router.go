package router

import (
    "github.com/gin-gonic/gin"
    "taskmanager/controllers"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()
    router.GET("/tasks", controllers.GetTasks)
    router.GET("/tasks/:id", controllers.GetTask)
	router.PUT("/tasks/:id",controllers.UpdateTask)
	router.DELETE("/tasks/:id",controllers.DeleteTask)
	router.POST("/tasks",controllers.Addtask)
    return router
}