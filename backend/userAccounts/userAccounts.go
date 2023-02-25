// package useraccounts
package userAccounts

// code from https://github.com/Duomly/go-bank-backend

import (
	"main/database"
	"main/interfaces"
)

// Refactor function updateAccount to use database package
func UpdateAccount(id uint, amount int) interfaces.ResponseAccount {
	account := interfaces.Account{}
	responseAcc := interfaces.ResponseAccount{}

	database.DB.Where("id = ? ", id).First(&account)
	account.Balance = uint(amount)
	database.DB.Save(&account)

	responseAcc.ID = account.ID
	responseAcc.Name = account.Name
	responseAcc.Balance = int(account.Balance)
	return responseAcc
}

// Refactor function getAccount to use database package
func getAccount(id uint) *interfaces.Account {

	account := &interfaces.Account{}
	if err := database.DB.Where("id = ? ", id).First(&account).Error; err != nil {
		return nil
	}
	return account
}
