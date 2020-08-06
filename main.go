package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/auth"
	"github.com/saedyousef/abwaab-task/controllers"
)

var ( 
	router = gin.Default()
)

func main() {
	// Connect to DB
	models.ConnectDatabase()

	// No authentication is required.
	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.CreateUser)
	router.POST("/refresh", auth.Refresh)
	
	// Authentication required.
	router.POST("/tweets/create", auth.TokenAuthMiddleware(), controllers.CreateTweet)
	router.GET("/twitter/search", auth.TokenAuthMiddleware(), controllers.SearchTweets)
	router.GET("/tweets", controllers.GetUserTweets)
	router.GET("/tweets/:tweetid", controllers.TweetDetails)
	router.POST("/tweets/update/:tweetid", controllers.UpdateTweet)
	router.DELETE("/tweets/:tweetid/delete", controllers.DeleteTweet)
	
	log.Fatal(router.Run(""))
}


