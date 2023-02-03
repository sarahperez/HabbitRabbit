package migrations

import (
	"example.com/backend_banking_app/helpers"
	"example.com/backend_banking_app/interfaces"
)

func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Username: "Martin", Email: "martin@martin.com"},
		{Username: "Michael", Email: "michael@martin.com"},
	}

	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
		user := &interfaces.User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)

		account := &interfaces.Account{Type: "Daily Account", Name: string(users[i].Username + "'s" + " account"), Balance: uint(10000 + int(i+1)), UserID: user.ID}
		db.Create(&account)
	}
	defer db.Closed()
}

func Migrate() {
	User := &interfaces.User{}
	Account := &interfaces.User{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User{}, &Account{})
	defer db.Closed()

	createAccounts()
}
