package main

import (
	"flag"
	"github.com/xercoy/glob"
)

func main() {
	var unit, mode, dest string
	var amount int

	flag.StringVar(&unit, "unit", "MB", "Unit of space for the glob.")
	flag.StringVar(&mode, "mode", "default", "Mode of the file content.")
	flag.IntVar(&amount, "amount", 1, "Number of files to be created.")
	flag.StringVar(&dest, "dest", "./", "Destination of created globs.")

	flag.Parse()

	g := glob.NewGlob(unit, mode, amount, dest)

	err := g.Make()
	if err != nil {
		panic(err.Error())
	}
}
