package main

import (
	"flag"
	"fmt"
	"github.com/xercoy/blobs"
	"log"
	"os"
)

func main() {
	var unit, dest, fmtStr, inputType string
	var amount int
	var random, help bool

	flag.StringVar(&unit, "unit", "1MB", "Unit of space for the glob.")
	flag.IntVar(&amount, "amount", 1, "Number of files to be created.")
	flag.StringVar(&dest, "dest", "./", "Destination of created globs.")
	flag.StringVar(&fmtStr, "o", "%d.dat", "Format specifier for blob file name. %d is for the number sequence of the file.")
	flag.BoolVar(&random, "random", false, "Random number of blobs ranging from 1 to the value of the amount flag")
	flag.BoolVar(&help, "help", false, "Displays flag attributes & usage information.")
	flag.StringVar(&inputType, "input-type", "zero", "Specifies input type of blob content. stdin = stdin, random = random characters, zero = zero characters.")

	flag.Parse()
	//	flag.CommandLine.SetOutput(os.Stdout)

	if help {
		fmt.Println("Blobs 1.0\n\n Usage: blobs <options> <options>...")
		flag.PrintDefaults()
		return
	}

	runner := blobs.NewRunner(os.Stdin, dest, unit, fmtStr, amount, random, inputType)

	err := blobs.Mk(runner)
	if err != nil {
		log.Fatalln(err)
	}
}
