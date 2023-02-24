// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package helpers

import (
	"testing"
)

func TestHelpersTest(t *testing.T) {
	expectedString := "helpers test is working"
	actualString := helpersTest()

	if expectedString != actualString {
		t.Errorf("Expected String(%s) is not same as"+
			" actual String (%s)", expectedString, actualString)

	}

}
