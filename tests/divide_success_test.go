package sampletest

import (
	"testing"
)

func TestDivide(t *testing.T) {

	result := divide(60, 2)

	if result < 1 {
		t.Error("Unexpected result. You cannot divide by zero")
	} else {
		t.Log("Successfully divided")
	}
}
