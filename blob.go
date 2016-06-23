package glob

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

// Runner contains information related to creating blobs.
type Runner struct {
	Src       io.Reader
	Dest      string
	Amount    int
	Unit      string
	FormatStr string
	Content   []byte
}

func NewRunner(src io.Reader, dst, unit, fmtStr string, amnt int) *Runner {
	r := new(Runner)

	r.Src = src
	r.Dest = dst
	r.Amount = amnt
	r.Unit = unit
	r.FormatStr = fmtStr

	return r
}

func (r *Runner) Mk() error {
	// Error handling and detection should be done here.

	srcContent, err := ioutil.ReadAll(r.Src)
	if err != nil {
		return err
	}
	r.Content = srcContent

	for i := 1; i <= r.Amount; i++ {
		log.Printf("Starting file %d...\n", i)

		err := r.createBlob(fmt.Sprintf(r.FormatStr, i))
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Runner) createBlob(fileName string) error {

	// Create a new file.
	fullBlobPath := filepath.Join(r.Dest, fileName)

	newFile, err := createBlobFile(fullBlobPath)
	if err != nil {
		return err
	}
	defer newFile.Close()

	// Fill it up!
	err = r.fillFile(newFile)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) fillFile(file *os.File) error {
	remainingBytes := BytesInUnit(r.Unit)
	var contentSize = len(r.Content)
	var endIndex int

	for {
		// Exit condition
		if remainingBytes == 0 {
			break
		}

		// If content is lesser in size than remaining bytes,
		// write all of the content to the file..
		if contentSize <= remainingBytes {
			endIndex = contentSize

			// If the content is larger in size than remaining bytes,
			// ensure that only the length of the remaining bytes are
			// copied from content.
		} else if contentSize >= remainingBytes {
			endIndex = remainingBytes
		}

		// Write the content to the file, subtract the number of
		// bytes written from the total of remainingBytes.
		bytesWritten, err := file.Write(r.Content[0:endIndex])
		if err != nil {
			return err
		}

		remainingBytes -= bytesWritten
	}

	return nil
}
