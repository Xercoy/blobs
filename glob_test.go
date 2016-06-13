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
