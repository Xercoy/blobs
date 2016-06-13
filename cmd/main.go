package main

import (
	"flags"
	"fmt"
	"github.com/xercoy/glob"
)

func main() {
	var unit, mode, dest string
	var amount int

	flags.StringVar(&unit, "MB", "Unit of space for the glob.")
	flags.StringVar(&mode, "default", "Mode of the file content.")
	flags.IntVar(&amount, 1, "Number of files to be created.")
	flags.StringVar(&dest, "./", "Destination of created globs.")

	g := glob.NewGlob(unit, mode, amount, dest)

	err := g.Make()
	if err != nil {
		panic(err.Error())
	}
}
