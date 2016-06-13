package glob

type Glob struct {
	Unit   string
	Amount int
	Mode   string
	Dest   string
}

func NewGlob(unit string, mode string, amt int, dest string) *Glob {
	g := new(Glob)

	g.Unit = unit
	g.Mode = mode
	g.Amount = amt
	g.Dest = dest

	return g
}
