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
