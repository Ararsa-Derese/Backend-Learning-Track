package controllers_test

import (
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestTaskController_GetTasks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.TaskUsecase)
	taskController := controllers.TaskController{TaskUsecase: mockUsecase}

	router := gin.Default()
	router.GET("/tasks", taskController.GetTasks)

	mockUsecase.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return([]domain.Task{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestTaskController_AddTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.TaskUsecase)
	taskController := controllers.TaskController{TaskUsecase: mockUsecase}

	router := gin.Default()

	router.POST("/tasks", taskController.AddTask)

	mockUsecase.On("AddTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("*domain.Task")).Return(nil)

	reqBody := strings.NewReader(`{"title": "New Task", "description": "Task Description"}`)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", reqBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestTaskController_UpdateTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.TaskUsecase)
	taskController := controllers.TaskController{TaskUsecase: mockUsecase}

	router := gin.Default()
	router.PUT("/tasks/:id", taskController.UpdateTask)

	mockUsecase.On("UpdateTask", mock.Anything, mock.AnythingOfType("*domain.Claims"),mock.AnythingOfType("primitive.ObjectID") ,mock.AnythingOfType("*domain.UpdateTask")).Return(nil)
	reqBody := strings.NewReader(`{"title": "New Task", "description": "Task Description"}`)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", reqBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestTaskController_DeleteTask(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.TaskUsecase)
	taskController := controllers.TaskController{TaskUsecase: mockUsecase}

	router := gin.Default()
	router.DELETE("/tasks/:id", taskController.DeleteTask)

	mockUsecase.On("DeleteTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}
