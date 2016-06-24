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

// NewRunner returns a pointer to a Runner and initializes most of the fields
// with the given args.
func NewRunner(src io.Reader, dst, unit, fmtStr string, amnt int) *Runner {
	r := new(Runner)

	r.Src = src
	r.Dest = dst
	r.Amount = amnt
	r.Unit = unit
	r.FormatStr = fmtStr

	return r
}

// Mk receives an instance of a Runner and creates blobs based on its attributes.
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

// createBlob creates the a blob with a given string and then writes to the file
// accordingly.
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

// fillBlob receives a pointer to a file which the function will write to.
// Specifications on what to write and how many bytes to write come from the
// Runner instance.
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

// createBlobFile creates a file at the specified path.
func createBlobFile(fullPath string) (*os.File, error) {
	return os.Create(fullPath)
}

// Unused, leaving this here because the humanized pkg is still new to me.
func bytesInUnit(unit string) int {
	var byteAmount int
	switch unit {
	case "B":
		byteAmount = 1
	case "KB":
		byteAmount = 1024
	case "MB":
		byteAmount = 1048576
	case "GB":
		byteAmount = 1073741824
	case "TB":
		byteAmount = 1099511627776
	}

	return byteAmount
}
