package glob

import (
	"os"
	"testing"
)

var (
	testDir = os.TempDir()
)

func TestMakeNewGlob(t *testing.T) {
	testGlob := NewGlob("MB", "default", 1, testDir)

	if testGlob == nil {
		t.Error("testGlob is nil, it shouldn't be.")
	}

	if testGlob.Unit != "MB" {

		t.Error("Glob field Unit did not match.")

	} else if testGlob.Mode != "default" {

		t.Error("Glob field Mode did not match.")

	} else if testGlob.Amount != 1 {

		t.Error("Glob field Amount did not match.")

	} else if testGlob.Dest != testDir {

		t.Error("Glob field Destination did not match.")

	}
}

func TestFieldValidation(t *testing.T) {
	var err error

	testGlob1 := NewGlob("B", "default", 1, testDir)
	testGlob2 := NewGlob("FOO", "BAR", 12309, testDir)

	err = testGlob1.FieldValidation()
	if err != nil {
		t.Errorf("Glob with valid fields determined invalid:\n%s",
			err.Error())
	}

	err = testGlob2.FieldValidation()
	if err == nil {
		t.Errorf("Glob with invalid fields determined invalid:\n%s",
			err.Error())
	}
}

func TestValidUnit(t *testing.T) {
	if validUnit("B") == false {
		t.Error("Valid unit value determined to be invalid.")
	}

	if validUnit("foobar") == true {
		t.Error("Invalid unit value deteremined to be valid.")
	}
}

func TestValidMode(t *testing.T) {
	if validMode("foobar") == true {
		t.Error("Invalid destination determined to be valid.")
	}

	if validMode("default") == false {
		t.Error("Valid destination determined to be invalid.")
	}
}