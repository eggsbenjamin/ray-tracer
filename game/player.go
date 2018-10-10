package game

import "math"

type Focusable interface {
	Pos() *Point
	Dir() float64
}

type Player struct {
	pos                         *Point
	dir, speed, rotatationSpeed float64
	states                      map[string]struct{}
	currentWeapon               Weapon
}

//	constructor
func NewPlayer(p *Point, dr, sp, rsp float64, w Weapon) *Player {
	return &Player{
		pos:             p,
		dir:             dr,
		speed:           sp,
		rotatationSpeed: rsp,
		states:          make(map[string]struct{}),
		currentWeapon:   w,
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
	p.states[v] = struct{}{}
}

func (p *Player) UnsetState(v string) {
	delete(p.states, v)
}

func (p *Player) Update(l *Level) {
	if _, ok := p.states["walking-forward"]; ok {
		n := GetEndPoint(p.Pos(), p.speed, p.Dir())
		if l.Walkable(n) {
			p.Move(n)
		}
	}
	if _, ok := p.states["walking-backward"]; ok {
		n := GetEndPoint(p.Pos(), -p.speed, p.Dir())
		if l.Walkable(n) {
			p.Move(n)
		}
	}
	if _, ok := p.states["strafing-right"]; ok {
		n := GetEndPoint(p.Pos(), p.speed, p.Dir()-math.Pi/2)
		if l.Walkable(n) {
			p.Move(n)
		}
	}
	if _, ok := p.states["strafing-left"]; ok {
		n := GetEndPoint(p.Pos(), -p.speed, p.Dir()+math.Pi/2)
		if l.Walkable(n) {
			p.Move(n)
		}
	}
	if _, ok := p.states["rotating-left"]; ok {
		p.Rotate(-p.rotatationSpeed)
	}
	if _, ok := p.states["rotating-right"]; ok {
		p.Rotate(p.rotatationSpeed)
	}
}
