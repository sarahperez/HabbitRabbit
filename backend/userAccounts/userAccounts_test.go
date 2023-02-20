// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package userAccounts

import (
	"testing"
)

func TestUserAccountsTest(t *testing.T) {
	expectedString := "userAccounts test is working"
	actualString := userAccountsTest()

	if expectedString != actualString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual String (%s)", expectedString, actualString)

	}

}
