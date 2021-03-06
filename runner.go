package blobs

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Runner contains information related to creating blobs.
type Runner struct {
	Src       io.Reader
	Dest      string
	Amount    int
	Unit      string
	FormatStr string
	Content   []byte
	Random    bool
	InputType string
}

// NewRunner returns a pointer to a Runner and initializes most of the fields
// with the given args.
func NewRunner(src io.Reader, dst, unit, fmtStr string, amnt int, random bool, inputType string) *Runner {
	r := new(Runner)

	r.Src = src
	r.Dest = dst
	r.Amount = amnt
	r.Unit = unit
	r.FormatStr = fmtStr
	r.Random = random
	r.InputType = inputType

	return r
}

// Mk receives an instance of a Runner and creates blobs based on its attributes.
func Mk(r *Runner) error {
	// Error handling and detection should be done here.
	err := r.validateFields()
	if err != nil {
		return err
	}

	// Fill content based on the mode and unit.
	/*	srcContent, err := ioutil.ReadAll(r.Src)
		if err != nil {
			return err
		}
		r.Content = srcContent*/

	srcContent, err := r.getContent()
	if err != nil {
		return err
	}
	r.Content = srcContent

	// Determine the number of blobs to be made. Takes the random flag into account.
	amount := r.setBlobAmount()

	log.Printf("Creating blobs.")
	startTime := time.Now()

	var wg sync.WaitGroup
	var doneChan = make(chan int)
	var errChan = make(chan error)

	for i := 1; i <= amount; i++ {
		wg.Add(1)

		fileName := fmt.Sprintf(r.FormatStr, i)
		index := i

		go func() {
			log.Printf("Creating file #%d of %d (%s)...\n", index, amount, fileName)
			err := r.createBlob(fileName)
			if err != nil {
				errChan <- err
			}

			defer wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case err = <-errChan:
		return err
	case <-doneChan:
	}

	endTime := time.Now()

	totalRuntime := endTime.Sub(startTime)
	log.Printf("Blob creation finished in %s seconds.", totalRuntime.String())

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

func (r *Runner) setBlobAmount() int {
	var amount int

	if r.Random == true {
		//randomizer := rand.New(rand.NewSource((int64)(amount)))
		seed := time.Now().UnixNano()

		log.Printf("Seed value is %v.\n", seed)

		randomizer := rand.New(rand.NewSource(seed))

		amount = randomizer.Intn(r.Amount)

		log.Printf("Random flag set, creating %d random files.\n", amount)
	} else {
		amount = r.Amount
	}

	return amount
}

func (r *Runner) getContent() ([]byte, error) {
	var content []byte

	switch r.InputType {
	case "zero":
		content = []byte{0}

	case "random":
		// Make a byte slize as large as the unit.
		byteAmount, err := humanize.ParseBigBytes(r.Unit)
		if err != nil {
			return nil, err
		}

		byteContent := make([]byte, (*byteAmount).Int64())

		randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))

		if _, err := randomizer.Read(byteContent); err != nil {
			return nil, err
		}

		content = byteContent

	case "stdin":
		byteContent, err := ioutil.ReadAll(r.Src)
		if err != nil {
			return nil, err
		}

		content = byteContent
	}

	return content, nil
}
