// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package users

import (
	"testing"
)

func TestPrepareToken(t *testing.T)    {}
func TestPrepareResponse(t *testing.T) {}
func TestLogin(t *testing.T)           {}
func TestRegister(t *testing.T)        {}
func TestGetUser(t *testing.T)         {}
