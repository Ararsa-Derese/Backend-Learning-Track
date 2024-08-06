package controllers

import (
	"net/http"
	"taskmanager/data"
	"taskmanager/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	// Implement user registration logic
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, errr := data.GetUserByID(user.ID)
	if errr == nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "User already exists"})
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	// adduser := models.User{
	// 	Username: user.Username,
	// 	Password: string(hashedPassword),
	// 	Role:     user.Role,
	// }
	adduser := *models.NewUser(user.Username, string(hashedPassword), user.Role)
	err = adduser.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = data.RegisterUser(&adduser)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"message": "Error registering the user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully", "userid": adduser.ID})

}
func LoginUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := data.GetUserByID(user.ID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := data.GenerateToken(result)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	c.Header("Authorization", "Bearer "+token) // Add token to the header
	c.JSON(http.StatusOK, gin.H{"Message": "User logged in successfully"})
}
