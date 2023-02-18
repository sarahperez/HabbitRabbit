package api

import (
	"testing"
)

func TestApiTest(t *testing.T) {
	expectedString := "api test is working"
	actualString := apiTest()

	if expectedString != actualString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual String (%s)", expectedString, actualString)

	}

}
