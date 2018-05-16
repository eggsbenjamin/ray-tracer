package game

type Focusable interface {
	Pos() *Point
	Dir() float64
}

type Player struct {
	pos   *Point
	dir   float64
	state string
}

//	constructor
func NewPlayer(p *Point, dr float64) *Player {
	return &Player{
		pos: p,
		dir: dr,
	}
}

func (p *Player) Pos() *Point {
	return p.pos
}

func (p *Player) Dir() float64 {
	return p.dir
}

func (p *Player) Rotate(v float64) {
	a := NormaliseAngle(p.dir + v)
	p.dir = a
}

func (p *Player) Move(n *Point) {
	p.pos.X = n.X
	p.pos.Y = n.Y
}

func (p *Player) SetState(v string) {
	p.state = v
}

func (p *Player) Update() {
	switch p.state {
	case "rotating":
		panic("TODO")
	}
}
