package game

import (
	"fmt"
	"math"
)

type Player struct {
	Pos *Point
	Dir float64
	Cam *Camera
}

//	constructor
func NewPlayer(x, y, dr float64) *Player {
	pos := NewPoint(x, y)
	return &Player{
		Pos: pos,
		Dir: dr,
		Cam: NewCamera(pos, dr, 0.5, 1),
	}
}

func (p *Player) Rotate(v float64) {
	p.Dir += v
	p.Cam.Angle += v
}

func (p *Player) Move(v float64) {
	n := GetEndPoint(p.Pos, v, p.Dir)
	p.Pos.X = n.X
	p.Pos.Y = n.Y
}
