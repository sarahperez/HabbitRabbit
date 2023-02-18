package users

import (
	"testing"
)

func TestUserTest(t *testing.T) {
	expectedString := "user test is working"
	actualString := userTest()

	if expectedString != actualString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual String (%s)", expectedString, actualString)

	}

}
