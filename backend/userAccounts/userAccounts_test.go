// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package userAccounts

import (
	"main/interfaces"
	"testing"
)

// TestUpdateAccount isnt running
//func TestUpdateAccount(t *testing.T) {
//expectedAccount := interfaces.Account{}
//expectedRespAcc := interfaces.ResponseAccount{}

//database.DB.Where("id = ? ", 1000).First(&expectedAccount)
//expectedAccount.Balance = uint(56)
//database.DB.Save(&expectedAccount)

//expectedRespAcc.ID = expectedAccount.ID
//expectedRespAcc.Name = expectedAccount.Name
//expectedRespAcc.Balance = int(expectedAccount.Balance)

//actualRespAcc := UpdateAccount(1000, 56)

//if expectedRespAcc != actualRespAcc {
//	t.Errorf("expected error to be nil got %v", actualRespAcc)
//}

//}

// TestGetAccount hasnt been started
func TestGetAccount(t *testing.T) {

}
