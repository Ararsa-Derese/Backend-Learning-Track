package controllers

import (
	"fmt"
	"net/http"
	"taskmanagerdb/data"
	"taskmanagerdb/models"
	"time"

	"github.com/gin-gonic/gin"
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
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	_, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.UpdateTask(id, &task)
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})

}
func Addtask(c *gin.Context) {
	var task models.Task
    
	// var adtask models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    _, errr := data.GetTaskByID(task.ID)
    if errr == nil {
        c.JSON(http.StatusNotAcceptable, gin.H{"message": "Task already exists"})
        return
    }
	duedate, er := time.Parse(time.RFC3339, task.DueDate)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format"})
		return
	}
	adtask := models.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     duedate.Format(time.RFC3339),
		Status:      task.Status,
	}
	err := adtask.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := data.AddTask(&adtask)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error adding the task"})
	}
	fmt.Println(&result)
	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully"})
}
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	_, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	data.DeleteTask(id)
	c.JSON(http.StatusAccepted, gin.H{"message": "Task Deleted successfully"})
}
