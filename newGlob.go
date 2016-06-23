package glob

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

var (
	once sync.Once
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
	srcContent, err := ioutil.ReadAll(r.Src)
	if err != nil {
		return err
	}

	for i := 1; i <= r.Amount; i++ {
		fileName := strconv.Itoa(i) + r.Unit

		err := CreateBlob(srcContent, r.Dest, fileName, r.Unit)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateBlob(content []byte, dst, fileName, unit string) error {
	fullBlobPath := filepath.Join(dst, fileName)

	// Create a new file.
	newFile, err := createBlobFile(fullBlobPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Fill it with junk.
	err = fillFile(newFile, content)
	if err != nil {
		return err
	}

	return nil
}

// should be the new create file, need to rename.
func createBlobFile(src string) (*os.File, error) {
	return os.Create(src)
}

func fillFile(file *os.File, content []byte) error {

	bytesWritten, err := file.Write(content)
	if err != nil {
		return err
	}

	log.Printf("%d bytes written to file.", bytesWritten)

	return err
}
