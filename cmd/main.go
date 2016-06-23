package main

import (
	"flag"
	"github.com/xercoy/blobs"
	"os"
)

func main() {
	var unit, dest, fmtStr string
	var amount int

	flag.StringVar(&unit, "unit", "MB", "Unit of space for the glob.")
	flag.IntVar(&amount, "amount", 1, "Number of files to be created.")
	flag.StringVar(&dest, "dest", "./", "Destination of created globs.")
	flag.StringVar(&fmtStr, "o", "%d.dat", "Format specifier for blobs.")

	flag.Parse()

	runner := blobs.NewRunner(os.Stdin, dest, unit, fmtStr, amount)

	err := blobs.Mk(runner)
	if err != nil {
		panic(err.Error())
	}
}
