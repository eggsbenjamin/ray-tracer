package game

var (
	BLUE  = NewColour(0, 92, 161, 1)
	BLACK = NewColour(0, 0, 0, 1)
)

type Palette map[int]*Colour

//	constructor
func NewPalette() Palette {
	return Palette{}
}

type Colour struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

//	constructor
func NewColour(r, g, b, a uint8) *Colour {
	return &Colour{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}
