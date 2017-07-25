package game

import "math"

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
	if p.Dir+v > 2*math.Pi {
		p.Dir = (p.Dir + v) - (2 * math.Pi)
		p.Cam.Angle = p.Dir
		return
	}
	if p.Dir+v < 0 {
		p.Dir = (2 * math.Pi) + (p.Dir + v)
		p.Cam.Angle = p.Dir
		return
	}
	p.Dir += v
	p.Cam.Angle = p.Dir
}

func (p *Player) Move(n *Point) {
	p.Pos.X = n.X
	p.Pos.Y = n.Y
}
