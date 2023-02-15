package main

import (
	"testing"
)

func TestTestingPractice(t *testing.T) {
	actualString := testingPractice()
	expectedString := "testing is working"

	if actualString != expectedString {
		t.Errorf("Expected Bool(%s) is not same as"+
			" actual Bool (%s)", expectedString, actualString)
	}

}
