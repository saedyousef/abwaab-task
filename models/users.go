// models/users.go

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Name 	  string `json:"author" binding:"required"`
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
  
	// Create book
	user := models.User{Username: input.Username, Password: input.Password}
	models.DB.Create(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": user})
}