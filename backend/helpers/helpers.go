package helpers

// this file contains code from https://github.com/Duomly/go-bank-backend/tree/Golang-course-Lesson-6

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"main/interfaces"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"

	"main/password"

	"golang.org/x/crypto/bcrypt"
)

// code from https://github.com/Duomly/go-bank-backend/tree/Golang-course-Lesson-6 below

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// function to hash passwords for security
func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

// Create panic handler
func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			error := recover()
			if error != nil {
				log.Println(error)

				resp := interfaces.ErrResponse{Message: "Internal server error"}
				json.NewEncoder(w).Encode(resp)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// token validation function
func ValidateToken(id string, jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	//parses the token with the token string, token object and key function
	//key function receives parsed but unverified token
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	HandleErr(err)
	//makes sure all parts of the token converted correctly
	var userId, _ = strconv.ParseFloat(id, 8)
	if token.Valid && tokenData["user_id"] == userId {
		return true
	} else {
		return false
	}
}

// --------------------------------------------our added functions-------------------------------------

// function to check if username is valid
func UsernameValidation(username string) bool {

	//https://github.com/usvc/go-password#usage
	customPolicy := password.Policy{
		MaximumLength:         32,
		MinimumLength:         6,
		MinimumLowercaseCount: 0,
		MinimumUppercaseCount: 0,
		MinimumNumericCount:   0,
		MinimumSpecialCount:   0,
		CustomSpecial:         []byte("~`!@#$%^&*()_-=+[{]}\\|;:'\"<>./?"),
	}

	if err := password.Validate(username, customPolicy); err != nil {
		return false
	}
	return true
}

func NameValidation(username string) bool {

	//https://github.com/usvc/go-password#usage
	customPolicy := password.Policy{
		MaximumLength:         32,
		MinimumLength:         2,
		MinimumLowercaseCount: 0,
		MinimumUppercaseCount: 0,
		MinimumNumericCount:   0,
		MinimumSpecialCount:   0,
		CustomSpecial:         []byte("~`!@#$%^&*()_-=+[{]}\\|;:'\"<>./?"),
	}

	if err := password.Validate(username, customPolicy); err != nil {
		return false
	}
	return true
}

// function to check if password is valid
func PasswordValidation(pass string) bool {
	//https://github.com/usvc/go-password#usage
	customPolicy := password.Policy{
		MaximumLength:         32,
		MinimumLength:         6,
		MinimumLowercaseCount: 1,
		MinimumUppercaseCount: 1,
		MinimumNumericCount:   1,
		MinimumSpecialCount:   1,
		CustomSpecial:         []byte("~`!@#$%^&*()_-=+[{]}\\|;:'\"<>./?"),
	}

	if err := password.Validate(pass, customPolicy); err != nil {
		return false
	}
	return true
}

// check to see if email is valid
func EmailValidation(email string) bool {
	//https://github.com/usvc/go-password#usage
	customPolicy := password.Policy{
		MaximumLength:         32,
		MinimumLength:         6,
		MinimumLowercaseCount: 0,
		MinimumUppercaseCount: 0,
		MinimumNumericCount:   0,
		MinimumSpecialCount:   2,
		CustomSpecial:         []byte(".@"),
	}

	if err := password.Validate(email, customPolicy); err != nil {
		return false
	}
	return true
}
