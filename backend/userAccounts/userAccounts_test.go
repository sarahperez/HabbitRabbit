// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package userAccounts

import (
	"testing"
)

// TestUpdateAccount isnt running
func TestUpdateAccount(t *testing.T) {
	if UpdateAccount(2, 500).Balance != 500 && UpdateAccount(2, 500).ID != 1 && UpdateAccount(2, 500).Name != "Michael" {
		t.Errorf("update account isnt working")
	}

}

// TestGetAccount isnt running
func TestGetAccount(t *testing.T) {
	if getAccount(1).Name != "Martin" {
		t.Errorf("Get account isnt working correctly")
	}
}
