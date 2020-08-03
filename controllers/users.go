package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/helper"
	"github.com/saedyousef/abwaab-task/auth"
)

type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
	Name 	  string `json:"name" binding:"required"`
}

type LoginInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
	
	var userObj models.User
	var count int
	models.DB.Where("username = ?", input.Username).Find(&userObj).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is already exists"})
		return
	}
	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirm password doesn't match"})
		return
	}
	// Create user
	hashedPwd := helper.hashAndSalt([]byte(input.Password))
	user := models.User{Username: input.Username, Password: hashedPwd, Name: input.Name}
	models.DB.Create(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var user models.User
	var input LoginInput

	if login := c.ShouldBindJSON(&input); login != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": login.Error()})
		return
	}
	
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": input.Password})
		return
	}
	
	
	if user.Username != input.Username || !helper.comparePasswords(user.Password, []byte(input.Password)) {
		c.JSON(http.StatusUnauthorized, "Please provide a valid credentials")
		return
	}
	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}
	c.JSON(http.StatusOK, tokens)
}

func TokenAuthMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
     err := auth.TokenValid(c.Request)
     if err != nil {
        c.JSON(http.StatusUnauthorized, err.Error())
        c.Abort()
        return
     }
     c.Next()
  }
}