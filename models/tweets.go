// models/users.go

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
)

type Tweet struct {
	ID   uint   `json:"id" gorm:"primary_key"`
	Body string `json:"body"`
}


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
	tweet := models.TWeet{Body: input.Body}
	models.DB.Create(&tweet)
  
	c.JSON(http.StatusOK, gin.H{"data": tweet})
}