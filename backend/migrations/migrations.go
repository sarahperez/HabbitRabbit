// package migrations
package migrations

// code from https://github.com/Duomly/go-bank-backend

import (
	"main/database"
	"main/helpers"
	"main/interfaces"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Refactor createAccounts to use database package
// parts of this arent valuable, but for the most part it is
func CreateAccounts() {
	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@michael.com"},
	}
	for i := 0; i < len(users); i++ {
		// Correct one way
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		database.DB.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 * int(i+1)), UserID: user.ID}
		database.DB.Create(&account)
	}
}

// Refactor Migrate
func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.Account{}
	database.DB.AutoMigrate(&User, &Account)

	CreateAccounts()
}
