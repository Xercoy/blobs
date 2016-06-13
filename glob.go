package glob

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type Glob struct {
	Unit   string
	Amount int
	Mode   string
	Dest   string
}

func (g *Glob) Make() error {

	for i := 1; i <= g.Amount; i++ {
		fileName := strconv.Itoa(i) + g.Unit

		err := createFile(g.Dest, fileName, g.Unit, g.Mode)
		if err != nil {
			return err
		}
	}

	return nil
}

func NewGlob(unit string, mode string, amt int, dest string) *Glob {
	g := new(Glob)

	g.Unit = unit
	g.Mode = mode
	g.Amount = amt
	g.Dest = dest

	return g
}

// Need to validate Destination.
func (g *Glob) FieldValidation() error {
	var errStr string

	if validUnit(g.Unit) == false {
		errStr += fmt.Sprintf("\nInvalid Unit value %s. Refer to help docs.",
			g.Unit)
	}

	if validMode(g.Mode) == false {
		errStr += fmt.Sprintf("\nInvalid Destination value %s. Refer to help docs.",
			g.Mode)
	}

	if errStr != "" {
		return errors.New(errStr)
	}

	return nil
}

func createFile(path string, fileName string, unit string, source string) error {
	err := os.Chdir(path)
	if err != nil {
		return err
	}

	cmd := exec.Command("touch", fileName)

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
