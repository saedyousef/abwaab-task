package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
)

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