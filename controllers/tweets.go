package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/tweeters"
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


// This api call a function that uses Twitter Client to get tweets.
func SearchTweets(c *gin.Context) {

	// Validate input
	url := c.Request.URL.Query()
	if url["query"] == nil || len(url["query"][0]) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "URL param 'query' is required"})
	  	return
	}

	tweets := tweeters.SearchTweets(url["query"][0])
	
	c.JSON(http.StatusOK, tweets)
}