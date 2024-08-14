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
	"github.com/stretchr/testify/suite"
)

type TaskControllerSuite struct {
	suite.Suite
	taskController *controllers.TaskController
	mockUsecase    *mocks.TaskUsecase
	router         *gin.Engine
}

func (suite *TaskControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(mocks.TaskUsecase)
	suite.taskController = &controllers.TaskController{TaskUsecase: suite.mockUsecase}
	suite.router = gin.Default()
	suite.router.GET("/tasks", suite.taskController.GetTasks)
	suite.router.POST("/tasks", suite.taskController.AddTask)
	suite.router.PUT("/tasks/:id", suite.taskController.UpdateTask)
	suite.router.DELETE("/tasks/:id", suite.taskController.DeleteTask)
}

func (suite *TaskControllerSuite) TearDownTest() {
	suite.mockUsecase.AssertExpectations(suite.T())
}


func (suite *TaskControllerSuite) TestGetTasks() {
	suite.mockUsecase.On("GetTasks", mock.Anything, mock.AnythingOfType("*domain.Claims")).Return([]domain.Task{}, nil)

	req, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
}

func (suite *TaskControllerSuite) TestAddTask() {
	suite.mockUsecase.On("AddTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("*domain.Task")).Return(nil)

	reqBody := strings.NewReader(`{"title": "New Task", "description": "Task Description"}`)
	req, _ := http.NewRequest(http.MethodPost, "/tasks", reqBody)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	
}

func (suite *TaskControllerSuite) TestUpdateTask() {
	suite.mockUsecase.On("UpdateTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID"), mock.AnythingOfType("*domain.UpdateTask")).Return(nil)

	reqBody := strings.NewReader(`{"title": "New Task", "description": "Task Description"}`)
	req, _ := http.NewRequest(http.MethodPut, "/tasks/1", reqBody)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
}

func (suite *TaskControllerSuite) TestDeleteTask() {
	suite.mockUsecase.On("DeleteTask", mock.Anything, mock.AnythingOfType("*domain.Claims"), mock.AnythingOfType("primitive.ObjectID")).Return(nil)

	req, _ := http.NewRequest(http.MethodDelete, "/tasks/1", nil)
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}
