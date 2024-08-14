package controllers_test

import (
	"bytes"
	"cleantaskmanager/delivery/controllers"
	"cleantaskmanager/domain"
	"cleantaskmanager/domain/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserControllerSuite struct {
	suite.Suite
	router       *gin.Engine
	userUsecase  *mocks.UserUsecase
	userController *controllers.UserController
}

func (suite *UserControllerSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.userUsecase = new(mocks.UserUsecase)
	suite.userController = &controllers.UserController{UserUsecase: suite.userUsecase}
	suite.router = gin.Default()
	suite.router.POST("/register", suite.userController.Register)
	suite.router.POST("/login", suite.userController.Login)
}

func (suite *UserControllerSuite) TearDownTest() {
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *UserControllerSuite) TestRegister() {
	user := domain.User{Username: "testuser", Password: "password"}
	suite.userUsecase.On("RegisterUser", mock.AnythingOfType("*domain.User")).Return(nil)
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(body))
	w := httptest.NewRecorder()

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

}

func (suite *UserControllerSuite) TestLogin() {
	user := domain.Login{ID: primitive.NewObjectID(), Password: "password"}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	suite.userUsecase.On("GetUserByID", user.ID).Return(&domain.User{Username: "testuser", Password: "password"}, nil)
	suite.userUsecase.On("Checkpassword", string(hashedPassword), "password").Return(nil)
	suite.userUsecase.On("GenerateToken", mock.AnythingOfType("*domain.User")).Return("string", nil)
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
