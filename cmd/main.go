package main

import (
	"flag"
	"github.com/xercoy/blobs"
	"os"
)

func main() {
	var unit, dest, fmtStr string
	var amount int
	var random bool

	flag.StringVar(&unit, "unit", "MB", "Unit of space for the glob.")
	flag.IntVar(&amount, "amount", 1, "Number of files to be created.")
	flag.StringVar(&dest, "dest", "./", "Destination of created globs.")
	flag.StringVar(&fmtStr, "o", "%d.dat", "Format specifier for blobs.")
	flag.BoolVar(&random, "random", false, "Random number of blobs, 1 - amount")

	flag.Parse()

	runner := blobs.NewRunner(os.Stdin, dest, unit, fmtStr, amount, random)

	err := blobs.Mk(runner)
	if err != nil {
		panic(err.Error())
	}
}
