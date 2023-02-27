// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package users

import (
	"main/interfaces"
	"testing"
)

// TestPrepareToken is not running yet
func TestPrepareToken(t *testing.T) {
	var testUser interfaces.User
	testUser.Username = "Martin"
	testUser.Email = "Martin@gmail.com"
	testUser.Password = "M@rtinsPw0rd"

	expectedString := "token"
	actualString := prepareToken(testUser)
	if expectedString != actualString {
		t.Errorf("prepareToken FAILED, expected -> %v, got -> %v", expectedString, actualString)
	} else {
		t.Errorf("prepareToken PASSED, expected -> %v, got -> %v", expectedString, actualString)
	}

}

// TestPrepareResponse hasnt been started
//func TestPrepareResponse(t *testing.T) {}

// TestLogin hasnt been started
//func TestLogin(t *testing.T) {}

// TestRegister hasnt been started
//func TestRegister(t *testing.T) {}

// TestGetUser hasnt been started
//func TestGetUser(t *testing.T) {}
