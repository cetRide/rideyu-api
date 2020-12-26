package users

import (
	"os"

	u "github.com/cetRide/rideyu-api/apihelpers"
	"github.com/cetRide/rideyu-api/helpers"
	"github.com/cetRide/rideyu-api/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccount(username, email, password string) map[string]interface{} {

	if username == "" {
		return u.Message(false, "Username field is empty")
	}

	if email == "" {
		return u.Message(false, "Email address field is empty")
	}

	if !helpers.ValidateEmail(email) {
		return u.Message(false, "Invalid email address.")
	}

	passwordValidationResponse := helpers.ValidatePassword(password)

	if !passwordValidationResponse["status"].(bool) {
		return u.Message(true, passwordValidationResponse["message"].(string))
	}
	user := &models.User{}

	models.GetDB().Where(&models.User{Email: email}).First(user)
	if user.Email != "" {
		return u.Message(false,
			"Email address already in use by another user.Try again using a new one")
	}

	models.GetDB().Where(&models.User{Username: username}).First(user)
	if user.Username != "" {
		return u.Message(false,
			"Username already in use by another user.Try again using a new one")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	saltedPassword := string(hashedPassword)

	if err := models.GetDB().Create(&models.User{
		Username: username,
		Email:    email,
		Password: saltedPassword}).
		Error; err != nil {
		return u.Message(false, "Unable to create user account")
	}
	tk := &helpers.Token{Username: username, Password: password}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	response := u.Message(true, "Your account is successfully created.")
	response["token"] = tokenString
	return response
}

func LoginUser(username string, password string) map[string]interface{} {
	user := &models.User{}
	err := models.GetDB().Where(&models.User{Username: username}).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Username not found")
		}
		return u.Message(false, "Connection error. Please retry")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return u.Message(false, "Incorrect password. Please try again")
	}

	//Create JWT token
	tk := &helpers.Token{Username: username, Password: password}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	response := u.Message(true, "Successful log in")
	response["token"] = tokenString
	response["user"] = user.ID
	return response
}
