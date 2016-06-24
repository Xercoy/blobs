package blobs

import (
	"fmt"
	"github.com/dustin/go-humanize"
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

func Mk(r *Runner) error {
	// Error handling and detection should be done here.

	srcContent, err := ioutil.ReadAll(r.Src)
	if err != nil {
		return err
	}
	r.Content = srcContent

	for i := 1; i <= r.Amount; i++ {
		fileName := fmt.Sprintf(r.FormatStr, i)
		log.Printf("Creating file #%d (%s)...\n", i, fileName)

		err := r.createBlob(fileName)
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
	err = r.fillBlob(newFile)
	if err != nil {
		return err
	}

	return nil
}

func (r *Runner) fillBlob(file *os.File) error {

	// Func returns the number of bytes of the string unit of type big.Int.
	bPtr, err := humanize.ParseBigBytes(r.Unit)
	if err != nil {
		return err
	}

	// Dereference the ptr, retrieve int64 value
	remainingBytes := (*bPtr).Int64()

	// Number of bytes of the content to be written.
	var contentSize = int64(len(r.Content))

	// Denotes the index of the end of the content to be written to.
	var endIndex int64

	for {
		// Exit condition, no more bytes to write.
		if remainingBytes == int64(0) {
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

		// Write the content to the file and subtract the number of
		// bytes written from the total of remainingBytes.
		bytesWritten, err := file.Write(r.Content[0:endIndex])
		if err != nil {
			return err
		}

		remainingBytes -= int64(bytesWritten)
	}

	return nil
}
