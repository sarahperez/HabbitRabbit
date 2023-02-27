package helpers

// code from https://github.com/Duomly/go-bank-backend

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

func HandleErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func HashAndSalt(pass []byte) string {
	hashed, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	HandleErr(err)

	return string(hashed)
}

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
		MinimumSpecialCount:   1,
		CustomSpecial:         []byte(".@"),
	}

	if err := password.Validate(email, customPolicy); err != nil {
		return false
	}
	return true
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

func ValidateToken(id string, jwtToken string) bool {
	cleanJWT := strings.Replace(jwtToken, "Bearer ", "", -1)
	tokenData := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(cleanJWT, tokenData, func(token *jwt.Token) (interface{}, error) {
		return []byte("TokenPassword"), nil
	})
	HandleErr(err)
	var userId, _ = strconv.ParseFloat(id, 8)
	if token.Valid && tokenData["user_id"] == userId {
		return true
	} else {
		return false
	}
}
