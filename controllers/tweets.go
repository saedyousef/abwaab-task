package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/biezhi/gorm-paginator/pagination"
	"github.com/saedyousef/abwaab-task/auth"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/tweeters"
)

type CreateTweetInput struct {
	Body  string `json:"body" binding:"required"`
	UserID uint64 `json:"user_id"`
}

type UpdateTweetInput struct {
	Body  string `json:"body" binding:"required"`
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

	// Extracting the token to get the userid.
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}

	for _, tweet := range tweets {
		go SaveTweet(tweet, tokenAuth.UserId)
	}

	c.JSON(http.StatusOK, gin.H{"tweets": tweets})
}

// This function will be called to save tweets returned from twitter search API.
func SaveTweet(body string, userId uint64) {
	tweet := models.Tweet{Body: body, UserID: userId}
	models.DB.Create(&tweet)
}

// Create a new Tweet
func CreateTweet(c *gin.Context) {
	var input CreateTweetInput

	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extracting the token to get the userid.
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
		return
	}
	
	// Preparing before save.
	tweet := models.Tweet{Body: input.Body, UserID: tokenAuth.UserId}

	// Save tweet.
	models.DB.Create(&tweet)

	c.JSON(http.StatusCreated, gin.H{"data": tweet})
}

// Getting tweets for logged in user.
func GetUserTweets(c *gin.Context) {

	// Extracting the token to get the userid.
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
		return
	}

	// Gettin page and limit param queries.
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	
	// Model.
	var tweets []models.Tweet
	
	// Paginiate the result.
    paginator := pagination.Paging(&pagination.Param{
        DB:      models.DB.Where("user_id =? ", tokenAuth.UserId),
        Page:    page,
        Limit:   limit,
        OrderBy: []string{"id desc"},
        ShowSQL: true,
	}, &tweets)
	
    c.JSON(http.StatusOK, paginator)
}

// Return single tweet.
func TweetDetails(c *gin.Context) {
	tweetid := c.Param("tweetid")

	var tweet models.Tweet
	models.DB.Where("id = ?", tweetid).First(&tweet)

	if tweet.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No data found."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tweet": tweet})
}

// Update user's tweet.
func UpdateTweet(c *gin.Context) {
	
	// Extracting the token to get the userid.
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
		return
	}

	// Getting tweetid from the url.
	tweetid := c.Param("tweetid")

	var tweet models.Tweet
	models.DB.Where("id = ?", tweetid).First(&tweet)

	if tokenAuth.UserId != tweet.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can not edit this tweet."})
		return
	}

	var input UpdateTweetInput
	// Validate input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet.Body = input.Body
	models.DB.Save(&tweet)

	c.JSON(http.StatusOK, gin.H{"tweet": tweet})
}

// Delete user's tweet.
func DeleteTweet(c *gin.Context) {

	// Extracting the token to get the userid.
	tokenAuth, err := auth.ExtractTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"unauthorized"})
		return
	}

	tweetid := c.Param("tweetid")

	var tweet models.Tweet
	models.DB.Where("id = ?", tweetid).First(&tweet)

	if tokenAuth.UserId != tweet.UserID {
		c.JSON(http.StatusForbidden, gin.H{"error": "You can not delete this tweet."})
		return
	}

	models.DB.Delete(&tweet)
	c.JSON(http.StatusNotFound, gin.H{"message": "tweet has been deleted."})
}