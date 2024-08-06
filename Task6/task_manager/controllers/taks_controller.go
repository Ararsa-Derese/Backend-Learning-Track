package controllers

import (
	"net/http"
	"taskmanager/data"
	"taskmanager/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	tasks := data.GetAllTasks(userID.(string), role.(string))
	if tasks == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tasks found"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
func GetTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")
	task, err := data.GetTaskByID(id, userID.(string), role.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}
func UpdateTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")
	_, er := data.GetTaskByID(id, userID.(string), role.(string))
	if er != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data.UpdateTask(id, userID.(string), role.(string), &task)
	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})

}
func Addtask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	var task models.Task

	// var adtask models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, errr := data.GetTaskByID(task.ID, userID.(string), role.(string))
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
		Title:       task.Title,
		Description: task.Description,
		DueDate:     duedate.Format(time.RFC3339),
		Status:      task.Status,
		UserID:      userID.(string),
	}

	adtask = *models.NewTask(userID.(string), task.Title, task.Description, duedate.Format(time.RFC3339), task.Status)
	err := adtask.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, errrr := data.AddTask(&adtask)
	if errrr != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error adding the task"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task added successfully", "taskid": adtask.ID})
}
func DeleteTask(c *gin.Context) {
	userID, _ := c.Get("user_id")
	role, _ := c.Get("role")
	id := c.Param("id")
	_, err := data.GetTaskByID(id, userID.(string), role.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	data.DeleteTask(id, userID.(string), role.(string))
	c.JSON(http.StatusAccepted, gin.H{"message": "Task Deleted successfully"})
}
