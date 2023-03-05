// package users
package users

// code from https://github.com/Duomly/go-bank-backend

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
func prepareResponse(user *interfaces.User, accounts []interfaces.ResponseAccount, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Accounts: accounts,
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
		// Find accounts for the user, parts of this code wont be necessary
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

		var response = prepareResponse(user, accounts, true)

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
func Register(username string, email string, pass string) map[string]interface{} {
	validUser := helpers.UsernameValidation(username)
	validPass := helpers.PasswordValidation(pass)
	validEmail := helpers.EmailValidation(email)

	if validUser && validPass && validEmail {
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Username: username, Email: email, Password: generatedPassword}

		existingUsername := database.DB.Where("username = ? ", username).First(&user).RowsAffected
		existingEmail := database.DB.Where("username = ? ", username).First(&user).RowsAffected

		if existingUsername > 0 && existingEmail == 0 {
			return map[string]interface{}{"message": "entered username is already taken"}
		} else if existingEmail > 0 && existingUsername == 0 {
			return map[string]interface{}{"message": "entered email is already associated with an account"}
		} else if existingUsername > 0 && existingEmail > 0 {
			return map[string]interface{}{"message": "entered username is taken and email is already associated with an acocunt"}
		}

		database.DB.Create(&user)
		//parts of this code wont be necessary, but for the most part its relevant
		account := &interfaces.Account{Type: "Daily Account", Name: string(username + "'s" + " account"), Balance: 0, UserID: user.ID}
		database.DB.Create(&account)

		accounts := []interfaces.ResponseAccount{}
		respAccount := interfaces.ResponseAccount{ID: account.ID, Name: account.Name, Balance: int(account.Balance)}
		accounts = append(accounts, respAccount)
		var response = prepareResponse(user, accounts, true)

		return response
	} else if !validUser && validPass && validEmail {
		return map[string]interface{}{"message": "entered username does not meet our requirements"}
	} else if validUser && !validPass && validEmail {
		return map[string]interface{}{"message": "entered password does not meet our requirements"}
	} else if validUser && validPass && !validEmail {
		return map[string]interface{}{"message": "entered email does not meet our requirements"}
	} else if !validUser && validPass && !validEmail {
		return map[string]interface{}{"message": "entered username and email do not meet our  requirements"}
	} else if !validUser && !validPass && validEmail {
		return map[string]interface{}{"message": "entered username and password do not meet our  requirements"}
	} else if validUser && !validPass && !validEmail {
		return map[string]interface{}{"message": "entered email and password do not meet our requirements"}
	} else {
		return map[string]interface{}{"message": "entered username, email and password do not meet our requirements"}
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
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ? ", user.ID).Scan(&accounts)

		var response = prepareResponse(user, accounts, false)
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}
}
