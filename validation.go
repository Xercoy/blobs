package blobs

import (
	"errors"
	"fmt"
	"github.com/dustin/go-humanize"
	"os"
)

// Boolean is taken care of by default.
// Dest = destination of blobs.
func (r *Runner) validateFields() error {
	if err := unitFieldInvalid(r.Unit); err != nil {
		return errors.New("Invalid unit field. Must contain a number and a valid data unit. Ex. `1MB`.")
	}

	if inputTypeFieldInvalid(r.InputType) {
		return errors.New("Invalid input-type field value. See usage info.")
	}

	if amountFieldInvalid(r.Amount) {
		return errors.New(fmt.Sprintf("amount field value invalid. Must be at least 1."))
	}

	if err := destFieldInvalid(r.Dest); err != nil {
		return errors.New(fmt.Sprintf("dest field invalid: %v", err))
	}

	if r.InputType == "stdin" {
		if stdinEmpty(r) {
			return errors.New("Nothing to read from Src.")
		}
	}

	return nil
}

/* Complain if there's more than one file being made with a formatStr that
doesn't include %d - otherwise it'd be writing to the same file name over and over
again... Maybe this is desired? Commenting out just in case.
func formatStrFieldInvalid(formatStr string, amount int) bool {
	validString := strings.Contains(formatStr, "%d")

	if !validString && amount > 1{
		return true
	}

	return false
}*/
func unitFieldInvalid(unit string) error {
	_, err := humanize.ParseBigBytes(unit)

	return err
}

func inputTypeFieldInvalid(inputType string) bool {
	switch inputType {
	case "random":
		return false
	case "zero":
		return false
	case "stdin":
		return false
	}

	return true
}

func amountFieldInvalid(amount int) bool {
	if amount <= 0 {
		return true
	}

	return false
}

// Ensures destination of blobs is valid.
func destFieldInvalid(dest string) error {
	_, err := os.Stat(dest)
	if err != nil {
		return err
	}

	return nil
}

func stdinEmpty(r *Runner) bool {
	return true
}
