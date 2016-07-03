package blobs

import (
	"fmt"

	//	"github.com/xercoy/blobs"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

var (
	tempDir = os.TempDir()
)

func TestCreateRandomAmount(t *testing.T) {
	log.Println("\n\nStarting TestCreateRandomAmount...")

	contentSrc := strings.NewReader("icecream")

	testRunner := NewRunner(contentSrc, tempDir, "1MB", "%d.dat", 15, true, "zero")

	err := Mk(testRunner)
	if err != nil {
		t.Error(err.Error())
	}
}

func TestNewRunner(t *testing.T) {
	log.Println("\n\nStarting TestNewRunner...")

	testReader := strings.NewReader("foobarbaz")
	testRunner := NewRunner(testReader, tempDir, "2MB", "%d.dat", 3, false, "zero")

	fmtString := "Given Runner field %s not equal to the given test value."
	var value string

	//Read from the Src field
	rdrContent, err := ioutil.ReadAll(testRunner.Src)
	if err != nil {
		t.Error(err.Error())
	}

	if testRunner.Amount != 3 {
		value = "Amount"
		t.Errorf(fmtString, value)

	} else if testRunner.Unit != "2MB" {
		value = "Unit"
		t.Errorf(fmtString, value)

	} else if testRunner.Dest != tempDir {
		value = "Dest"
		t.Errorf(fmtString, value)

	} else if testRunner.FormatStr != "%d.dat" {
		value = "FormatStr"
		t.Errorf(fmtString, value)

	} else if (string)(rdrContent) != "foobarbaz" {
		value = "Src"
		t.Errorf(fmtString, value)
	}
}

func TestMk(t *testing.T) {
	log.Println("\n\nStarting TestMk...")

	testReader := strings.NewReader("helloWorld")
	testRunner := NewRunner(testReader, tempDir, "2MB", "%d.dat", 5, false, "zero")

	err := Mk(testRunner)
	if err != nil {
		t.Error(err.Error())
	}

	// Doesn't work for some reason:
	// wildcardStr := strings.Replace(testRunner.FormatStr, "%d", "*", -1)
	// Iterate through and ls each file for now...
	for i := 1; i <= testRunner.Amount; i++ {
		fileName := fmt.Sprintf(testRunner.FormatStr, i)

		blobFilter := filepath.Join(os.TempDir(), fileName)

		cmd := exec.Command("ls", "-l", blobFilter)
		cmdOutput, err := cmd.CombinedOutput()
		if err != nil {
			t.Error(err.Error())
		}

		log.Printf("%s", cmdOutput)
	}
}
