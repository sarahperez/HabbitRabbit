package users

import (
	"time"

	"example.com/backend_banking_app/helpers"
	"example.com/backend_banking_app/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func Login(username string, pass string) map[string]interface{} {
	//creates database connection
	db := helpers.ConnectDB()
	user := &interfaces.User{}
	//checks to see if the username exists in the database
	if db.Where("username = ?", username).First(&user).RecordNotFound() {
		return map[string]interface{}{"message": "User not found"}
	}

	//verify password
	//password verification
	passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	//checks if the password is not mismatched
	if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
		return map[string]interface{}{"message": "Wrong password"}
	}
	//bank account for user
	accounts := []interfaces.ResponseAccount{}
	//assign data from the database
	db.Table("account").Select("id, name, balance").Where("user_id = ?", user.ID).Scan(&accounts)

	//Setup response
	responseUser := &interfaces.ResponseUser{
		ID:       user.ID,
		Email:    user.Email,
		Accounts: accounts,
	}

	defer db.Close()

	//Sign token
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	//Prepare Response
	var response = map[string]interface{}{"message": "all is fine"}
	response["jwt"] = token
	response["data"] = responseUser

	return response

}
