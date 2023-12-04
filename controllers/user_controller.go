package controllers

import (
	"net/http"
	"rakamin-golang/database"
	"rakamin-golang/helpers"
	"rakamin-golang/models"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// User registration
func (uc *UserController) Register(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := govalidator.ValidateStruct(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, _ := helpers.HashPassword(userInput.Password)
	userInput.Password = hashedPassword

	database.DB.Create(&userInput)

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// User login
func (uc *UserController) Login(c *gin.Context) {
	var userInput models.User

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	database.DB.Where("email = ?", userInput.Email).First(&user)

	if user.ID == 0 || !helpers.CheckPassword(userInput.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Update user profile
func (uc *UserController) UpdateUser(c *gin.Context) {
	// userID := helpers.ExtractUserID(c)
	userID := c.Param("userId")
	var updatedUser models.User

	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	database.DB.Where("id = ?", userID).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Update user fields based on your requirements
	user.Username = updatedUser.Username
	database.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// Delete user account
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("userId")

	user := models.User{}
	database.DB.Where("id = ?", userID).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	database.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
