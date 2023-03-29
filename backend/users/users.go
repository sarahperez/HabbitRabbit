// package users
package users

// code from https://github.com/Duomly/go-bank-backend
//fixed merge branch?

import (
	"time"

	"main/database"
	"main/helpers"
	"main/interfaces"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Refactor prepareToken
// part of login function
// if the current time is after the time its est to expire at, the token has expired
func prepareToken(user *interfaces.User) string {
	//creates token using the user id and the current time as the duration
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute * 60).Unix(),
	}
	//creates a new jwt token using the signing method HS256 and the claims
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	//generates the token string
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token
}

// Refactor prepareResponse
func prepareResponse(user *interfaces.User, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}
	var response = map[string]interface{}{"message": "all is fine"}
	// Add withToken feature to prepare response
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser
	return response
}

// Refactor Login function to use database package
func Login(username string, pass string) map[string]interface{} {
	// Add validation to login
	validUser := helpers.UsernameValidation(username)
	validPass := helpers.PasswordValidation(pass)

	if validUser && validPass {
		user := &interfaces.User{}
		if err := database.DB.Where("username = ? ", username).First(&user).Error; err != nil {
			return map[string]interface{}{"message": "User not found"}
		}
		// Verify password
		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}

		var response = prepareResponse(user, true)

		return response
	} else if !validUser && validPass {
		return map[string]interface{}{"message": "entered username does not meet our requirements"}
	} else if validUser && !validPass {
		return map[string]interface{}{"message": "entered password does not meet our requirements"}
	} else {
		return map[string]interface{}{"message": "entered username and password do not meet our requirements"}
	}
}

// Refactor Register function to use database package
func Register(username string, name string, email string, pass string) map[string]interface{} {
	validUser := helpers.UsernameValidation(username)
	validName := helpers.NameValidation(name)
	validPass := helpers.PasswordValidation(pass)
	validEmail := helpers.EmailValidation(email)

	if validUser && validName && validPass && validEmail {
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Name: name, Email: email, Password: generatedPassword}

		var existingUsername int64
		var existingEmail int64
		database.DB.Where("username = ? ", username).Count(&existingUsername)
		database.DB.Where("email = ? ", email).Count(&existingEmail)

		if existingUsername > 0 && existingEmail == 0 {
			return map[string]interface{}{"message": "entered username is already taken"}
		} else if existingEmail > 0 && existingUsername == 0 {
			return map[string]interface{}{"message": "entered email is already associated with an account"}
		} else if existingUsername > 0 && existingEmail > 0 {
			return map[string]interface{}{"message": "entered username is taken and email is already associated with an acocunt"}
		}

		database.DB.Create(&user)

		var response = prepareResponse(user, true)

		return response
	} else if !validUser && validPass && validEmail && validName {
		return map[string]interface{}{"message": "entered username does not meet our requirements"}
	} else if validUser && !validPass && validEmail && validName {
		return map[string]interface{}{"message": "entered password does not meet our requirements"}
	} else if validUser && validPass && !validEmail && validName {
		return map[string]interface{}{"message": "entered email does not meet our requirements"}
	} else if validUser && validPass && validEmail && !validName {
		return map[string]interface{}{"message": "entered name does not meet our requirements"}
	} else if !validUser && validPass && !validEmail && validName {
		return map[string]interface{}{"message": "entered username and email do not meet our  requirements"}
	} else if !validUser && !validPass && validEmail && validName {
		return map[string]interface{}{"message": "entered username and password do not meet our  requirements"}
	} else if validUser && !validPass && !validEmail && validName {
		return map[string]interface{}{"message": "entered email and password do not meet our requirements"}
	} else if !validUser && validPass && validEmail && !validName {
		return map[string]interface{}{"message": "entered username and name do not meet our  requirements"}
	} else if validUser && !validPass && validEmail && !validName {
		return map[string]interface{}{"message": "entered name and password do not meet our  requirements"}
	} else if validUser && validPass && !validEmail && !validName {
		return map[string]interface{}{"message": "entered email and name do not meet our requirements"}
	} else if !validUser && validPass && !validEmail && !validName {
		return map[string]interface{}{"message": "entered username, name and email do not meet our  requirements"}
	} else if !validUser && !validPass && validEmail && !validName {
		return map[string]interface{}{"message": "entered username, name and password do not meet our  requirements"}
	} else if validUser && !validPass && !validEmail && !validName {
		return map[string]interface{}{"message": "entered email, name and password do not meet our requirements"}
	} else if validUser && !validPass && !validEmail && validName {
		return map[string]interface{}{"message": "entered username, email and password do not meet our requirements"}
	} else {
		return map[string]interface{}{"message": "entered fields do not meet our requirements"}
	}
}

// Refactor GetUser function to use database package
func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)
	// Find and return user
	if isValid {
		user := &interfaces.User{}
		if err := database.DB.Where("id = ? ", id).First(&user).Error; err != nil {
			return map[string]interface{}{"message": "User not found"}
		}

		var response = prepareResponse(user, false)
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
