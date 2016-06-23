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
		log.Printf("Starting file %d...\n", i)

		fileName := strconv.Itoa(i) + r.Unit

		err := CreateBlob(srcContent, r.Dest, fileName, r.Unit, r.Amount)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateBlob(content []byte, dst, fileName, unit string, amount int) error {
	fullBlobPath := filepath.Join(dst, fileName)

	// Create a new file.
	newFile, err := createBlobFile(fullBlobPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Fill it with junk.
	err = fillFile(newFile, content, unit, amount)
	if err != nil {
		return err
	}

	return nil
}

// should be the new create file, need to rename.
func createBlobFile(src string) (*os.File, error) {
	return os.Create(src)
}

/* Fill the file up to the specified amount.

Cases:
- buffer less than amount? Repeat
- buffer more than amount? Truncate */
func fillFile(file *os.File, content []byte, unit string, amount int) error {

	// Figure out how much to write, then write it.
	remainingBytes := BytesInUnit(unit)
	for {
		if remainingBytes == 0 {
			break
		}

		var bytesToBeWritten []byte
		var contentSize = len(content)

		if contentSize <= remainingBytes {
			bytesToBeWritten = content[0:contentSize]
		} else if contentSize >= remainingBytes {
			bytesToBeWritten = content[0:remainingBytes]
		}

		//write
		bytesWritten, err := file.Write(bytesToBeWritten)
		if err != nil {
			return err
		}

		remainingBytes = remainingBytes - bytesWritten
	}
	/*
		bytesWritten, err := file.Write(content)
		if err != nil {
			return err
		}

		log.Printf("%d bytes written to file.", bytesWritten)

		return err*/

	return nil
}
