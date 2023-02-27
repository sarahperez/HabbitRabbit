// to perform tests:
//cd into file you want to test, and type go test
//if multiple functions in test file, do go test -v instead of go test

package helpers

import (
	"testing"
)

// TestHandleErrTest hasnt been started
func TestHandleErrTest(t *testing.T) {}

// TestHashAndSaltTest hasnt been started
func TestHashAndSaltTest(t *testing.T) {}

// TestUsernameValidation is working
func TestUsernameValidation(t *testing.T) {
	if UsernameValidation("bad") == true {
		t.Errorf("username validation isnt working, short username returned true")
	}
	if UsernameValidation("testingreallyreallyreallylongusername") == true {
		t.Errorf("username validation isnt working, long username returned true")
	}
}

// TestPasswordValidation is working
func TestPasswordValidation(t *testing.T) {
	if PasswordValidation("B@d1") == true {
		t.Errorf("password validation isnt working, short password returned true")
	}
	if PasswordValidation("Testingre@1lyreallyreallylongusername") == true {
		t.Errorf("password validation isnt working, long password returned true")
	}
	if PasswordValidation("n0upperc@se") == true {
		t.Errorf("password validation isnt working, password with no uppercase returned true")
	}
	if PasswordValidation("NOL0WERC@SE") == true {
		t.Errorf("password validation isnt working, password with no lowercase returned true")
	}
	if PasswordValidation("noNuminP@ss") == true {
		t.Errorf("password validation isnt working, password with no numbers returned true")
	}
	if PasswordValidation("noSpecia1char") == true {
		t.Errorf("password validation isnt working, password with no special characters returned true")
	}

}

// TestEmailValidation is working
func TestEmailValidation(t *testing.T) {}

// TestPanicHandler hasnt been started
func TestPanicHandler(t *testing.T) {}

// TestValidateToken hasnt been started
func TestValidateToken(t *testing.T) {}
