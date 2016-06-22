package glob

import (
	"io"
	"os/exec"
	"path/filepath"
	"strconv"
)

// Runner contains information related to creating blobs.
type Runner struct {
	Src    io.Reader
	Dest   string
	Amount int
	Unit   string
}

func NewRunner(src io.Reader, dst, unit string, amnt int) *Runner {
	r := new(Runner)

	r.Src = src
	r.Dest = dst
	r.Amount = amnt
	r.Unit = unit

	return r
}

func (r *Runner) Mk() error {
	for i := 1; i <= r.Amount; i++ {
		fileName := strconv.Itoa(i) + r.Unit

		err := newCreateFile(r.Src, r.Dest, fileName, r.Unit)
		if err != nil {
			return err
		}
	}

	return nil
}

func newCreateFile(src io.Reader, dst, fileName, unit string) error {
	cmd := exec.Command("touch", filepath.Join(dst, fileName))

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
