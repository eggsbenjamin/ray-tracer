package game

type Player struct {
	Pos *Point
	Dir float64
	Cam *Camera
}

//	constructor
func NewPlayer(m *Map, x, y, dr float64) *Player {
	pos := NewPoint(x, y)
	return &Player{
		Pos: pos,
		Dir: dr,
		Cam: NewCamera(pos, m, dr, 0.5, 0.75),
	}
}

func (p *Player) Rotate(v float64) {
	a := NormaliseAngle(p.Dir + v)
	p.Dir = a
	p.Cam.Angle = p.Dir
}

func (p *Player) Move(n *Point) {
	p.Pos.X = n.X
	p.Pos.Y = n.Y
}
