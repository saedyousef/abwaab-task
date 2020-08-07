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
	router.POST("/auth/login", controllers.Login)
	router.POST("/auth/signup", controllers.CreateUser)
	router.POST("/auth/refresh", auth.Refresh)
	router.GET("/tweets/:tweetid", controllers.TweetDetails)
	
	// Authentication required.
	router.POST("/tweets/create", auth.TokenAuthMiddleware(), controllers.CreateTweet)
	router.GET("/twitter/search", auth.TokenAuthMiddleware(), controllers.SearchTweets)
	router.GET("/tweets", controllers.GetUserTweets)
	router.PUT("/tweets/:tweetid/update", auth.TokenAuthMiddleware(), controllers.UpdateTweet)
	router.DELETE("/tweets/:tweetid/delete", auth.TokenAuthMiddleware(), controllers.DeleteTweet)
	
	
	log.Fatal(router.Run(""))
}


