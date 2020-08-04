package controllers

import (
	"net/http"
	// "encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/tweeters"
)

type CreateTweetInput struct {
	Body  string `json:"body" binding:"required"`
}

type SearchTweetsInput struct {
	Query string `json:"query" binfing:"required"` 
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

func SearchTweets(c *gin.Context) {
	// Validate input
	var input SearchTweetsInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
	
	tweets := tweeters.SearchTweets(input.Query)
	// var response map[string]interface{}
	// json.Unmarshal([]byte(tweets), &response)
  
	c.JSON(http.StatusOK, tweets)
}