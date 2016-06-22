package glob

import (
	"os"
	"testing"
)

func TestNewRunner(t *testing.T) {
	testRunner := NewRunner(os.Stdin, "./", "MB", 5)

	if testRunner.Amount != 5 {
		t.Error("Given amount not equal to actual amount.")
	} else if testRunner.Unit != "MB" {
		t.Error("Given unit not equal to actual unit.")
	}
}
