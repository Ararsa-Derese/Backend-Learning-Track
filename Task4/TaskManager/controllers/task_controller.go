package controllers

import (
	"taskmanager/data"
	"taskmanager/models"

	"github.com/gin-gonic/gin"

	// "taskmanager/models"
	"net/http"
)

func GetTasks(c *gin.Context) {
    tasks := data.GetAllTasks()
    c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {
    id := c.Param("id")
    task, err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }
    c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context){
    id:= c.Param("id")
    _,err := data.GetTaskByID(id)
     if err != nil {
        c.JSON(http.StatusNotFound,gin.H{"error":"Task not found"})
        return
    }
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
   
    data.UpdateTask(id,&task)
    c.JSON(http.StatusOK,gin.H{"message":"Task updated successfully"})
    
    
}
func Addtask(c *gin.Context){
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    data.AddTask(&task)
    c.JSON(http.StatusOK,gin.H{"message":"Task added successfully"})
}
func DeleteTask(c *gin.Context){
    id := c.Param("id")
    _,err := data.GetTaskByID(id)
    if err != nil {
        c.JSON(http.StatusNotFound,gin.H{"error":"Task not found"})
        return
    }
    data.DeleteTask(id)
    c.JSON(http.StatusAccepted, gin.H{"message":"Task Deleted successfully"})
}

// Implement AddTask, UpdateTask, DeleteTask