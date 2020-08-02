package controllers

import (
	"net/http"
	"os"
	"time"
	"log"
	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
)

type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
	Name 	  string `json:"name" binding:"required"`
}

type LoginInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
}

func CreateUser(c *gin.Context) {
	// Validate input
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
	  	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  	return
	}
	
	var userObj models.User
	var count int
	models.DB.Where("username = ?", input.Username).Find(&userObj).Count(&count)
	if count != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is already exists"})
		return
	}
	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirm password doesn't match"})
		return
	}
	// Create user
	hashedPwd := hashAndSalt([]byte(input.Password))
	user := models.User{Username: input.Username, Password: hashedPwd, Name: input.Name}
	models.DB.Create(&user)
  
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func Login(c *gin.Context) {
	var user models.User
	var input LoginInput

	if login := c.ShouldBindJSON(&input); login != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": login.Error()})
		return
	}
	
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": input.Password})
		return
	}
	
	
	if user.Username != input.Username || !comparePasswords(user.Password, []byte(input.Password)) {
		c.JSON(http.StatusUnauthorized, "Please provide a valid credentials")
		return
	}
	token, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
}

func CreateToken(userId uint) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}