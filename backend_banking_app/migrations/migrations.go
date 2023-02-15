package migrations

import (
	"example.com/backend_banking_app/database"
	"example.com/backend_banking_app/helpers"
	"example.com/backend_banking_app/interfaces"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Refactor createAccounts to use database package
//parts of this arent valuable, but for the most part it is
func createAccounts() {
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
	//this can be deleted
	Transactions := &interfaces.Transaction{}
	//
	database.DB.AutoMigrate(&User, &Account, &Transactions)
	
	createAccounts()
}
// Delete Migrate transactions