package helpers

// code from https://github.com/Duomly/go-bank-backend

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"main/interfaces"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/lib/pq"
	"github.com/usvc/go-password"
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

// Create validation
func Validation(values []interfaces.Validation) bool {
	//checks to make sure username and id are valid
	username := regexp.MustCompile(`^([A-Za-z0-9]{5,})+$`)
	email := regexp.MustCompile(`^[A-Za-z0-9]+[@]+[A-Za-z0-9]+[.]+[A-Za-z]+$`)

	for i := 0; i < len(values); i++ {
		switch values[i].Valid {
		case "username":
			if !username.MatchString(values[i].Value) {
				return false
			}
		case "email":
			if !email.MatchString(values[i].Value) {
				return false
			}
		case "password":
			if len(values[i].Value) < 5 {
				return false
			}

			//------------------------------------------our added password requirements-------------------------------------
			//https://github.com/usvc/go-password#usage
			customPolicy := password.Policy{
				MaximumLength:         32,
				MinimumLength:         12,
				MinimumLowercaseCount: 1,
				MinimumUppercaseCount: 1,
				MinimumNumericCount:   1,
				MinimumSpecialCount:   1,
				CustomSpecial:         []byte("`!@"),
			}

			if err := password.Validate(values[i].Value, customPolicy); err != nil {
				log.Print("password is invalid")
				return false
			}
			//--------------------------------------------------------------------------------------------------------------
		}
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

func helpersTest() string {
	return "helpers test is working"
}
