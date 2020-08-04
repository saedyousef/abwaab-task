package controllers

import (
	"log"
	"time"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/saedyousef/abwaab-task/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/saedyousef/abwaab-task/auth"
)

// User signup struct.
type CreateUserInput struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	PasswordConfirm  string `json:"password_confirm" binding:"required"`
	Name 	  string `json:"name" binding:"required"`
}

// User login struct.
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

	// Check if password and password_confirm matches.
	if input.Password != input.PasswordConfirm {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password and confirm password doesn't match"})
		return
	}
	// Create user
	hashedPwd := hashAndSalt([]byte(input.Password))
	user := models.User{Username: input.Username, Password: hashedPwd, Name: input.Name}
	models.DB.Create(&user)

	// Log the registerd user in.
	var jsonResponse map[string]interface{}
	response := loginUser(input.Username, input.Password)
	json.Unmarshal([]byte(response), &jsonResponse)

	c.JSON(http.StatusCreated, jsonResponse)
}

// Basic login function, returns access_token & refresh_token.
func Login(c *gin.Context) {
	var user models.User
	var input LoginInput

	if login := c.ShouldBindJSON(&input); login != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": login.Error()})
		return
	}
	
	if err := models.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	
	// Checking for credentials.
	if user.Username != input.Username || !comparePasswords(user.Password, []byte(input.Password)) {
		c.JSON(http.StatusUnauthorized, "Please provide a valid credentials")
		return
	}

	// Create the auth token.
	ts, err := auth.CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	// Mapping tokens.
	tokens := map[string]string{
		"access_token":  ts.AccessToken,
		"refresh_token": ts.RefreshToken,
	}

	c.JSON(http.StatusOK, tokens)
}

// Generate hashed password using salt key.
func hashAndSalt(pwd []byte) string {

	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

// Caomparing hashed password with a plain password.
func comparePasswords(hashedPwd string, plainPwd []byte) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// Login user, this function is created to login user after signup.
func loginUser(username string, password string) string{

	// Setting up the request timeout and initialize http client.
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// Setting up request body.
	requestBody, err := json.Marshal(map[string]string {
		"username": username,
		"password": password,
	}) 

	if err != nil {
		return "authentications failed"
	}

	// Preparing the request.
	request, err := http.NewRequest("POST", "http://127.0.0.1/login", bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return "Request failed"
	}

	// Requesting.
	response, err := client.Do(request)
	if err != nil {
		return "Login Request falied"
	}
	
	defer response.Body.Close()

	// Read the response body.
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "Response failure"
	}

	// Return the response body(access_token, refresh_token)
	return string(body)
}