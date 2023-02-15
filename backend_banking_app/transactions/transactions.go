package transactions
//most of this code is not necessary, however may be helpful when dealing with friends or other aspects of our account later down the line

import (
	"example.com/backend_banking_app/database"
	"example.com/backend_banking_app/helpers"
	"example.com/backend_banking_app/interfaces"
)
// Refactor CreateTransaction to use database package
func CreateTransaction(From uint, To uint, Amount int) {
	transaction := &interfaces.Transaction{From: From, To: To}
	database.DB.Create(&transaction)
}

//looks for transactions that are from or to the id
func GetTransactionsByAccount(id uint) []interfaces.ResponseTransaction{
	transactions := []interfaces.ResponseTransaction{}
	database.DB.Table("transactions").Select("id, transactions.from, transactions.to, amount").Where(interfaces.Transaction{From: id}).Or(interfaces.Transaction{To: id}).Scan(&transactions)
	return transactions
}

// Create function GetMyTransactions
//this function may be relevent for finding and connecting with friends, but the specifics arent relevant
func GetMyTransactions(id string, jwt string) map[string]interface{} {
	// Validate JWT token
	isValid := helpers.ValidateToken(id, jwt)
	if isValid {
		// Find and return transactions related to user, may be relevant for friends
		accounts := []interfaces.ResponseAccount{}
		database.DB.Table("accounts").Select("id, name, balance").Where("user_id = ? ", id).Scan(&accounts)
		//

		//iterates through accounts and looks for transactions related to account and puts in array
		transactions := []interfaces.ResponseTransaction{}
		for i := 0; i < len(accounts); i++ {
			accTransactions := GetTransactionsByAccount(accounts[i].ID)
			transactions = append(transactions, accTransactions...)
		}

		var response = map[string]interface{}{"message": "all is fine"}
		response["data"] = transactions
		return response
	} else {
		return map[string]interface{}{"message": "Not valid token"}
	 }
}