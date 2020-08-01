package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
)

type CreateTweetInput struct {
	Body  string `json:"body" binding:"required"`
}

func CreateTweet(c *gin.Context) {
	// Validate input
	var input CreateTweetInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
  
	// Create book
	tweet := models.Tweet{Body: input.Body}
	models.DB.Create(&tweet)
  
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}