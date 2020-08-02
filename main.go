package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"github.com/saedyousef/abwaab-task/controllers"
)

var (
	router = gin.Default()
)

func main() {

	// Connect to DB
	models.ConnectDatabase()

	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.CreateUser)
	router.POST("/test", controllers.CreateTweet)
	log.Fatal(router.Run(""))
}


