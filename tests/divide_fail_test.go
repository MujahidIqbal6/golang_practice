package sampletest

import (
	"testing"
)

func TestDivideFail(t *testing.T) {

	result := divide(6, 0)

	if result < 1 {
		t.Error("Unexpected result. You cannot divide by zero")
	}

}
