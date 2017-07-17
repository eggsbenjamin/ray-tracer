package game

import "math"

type Vector []*Point

//	constructor
func NewVector(p1, p2 Point) *Vector {
	v := Vector(DDA(p1, p2))
	return &v
}

//	get the initial point of the vector
func (v Vector) Start() *Point {
	return v[0]
}

//	get the terminal point of the vector
func (v Vector) End() *Point {
	return v[len(v)-1]
}

// get the magnitude (length) of the vector
func (v Vector) Magnitude() float64 {
	e := v.End()
	a := math.Pow(e.X, 2)
	b := math.Pow(e.Y, 2)
	return math.Sqrt(a + b)
}

//	calculates and returns, in radians, the angle of the vector
func (v Vector) Angle() float64 {
	s := v.Start()
	e := v.End()
	return math.Atan2(e.Y-s.Y, e.X-s.X)
}
