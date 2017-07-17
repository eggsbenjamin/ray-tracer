package game

import "math"

//	dda algorithm implemented to obtain array of points on a line where
//	the line intersects with a an integer on either the x or y axis
func DDA(p1, p2 *Point) (v []*Point) {
	var steps float64
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	if math.Abs(dx) > math.Abs(dy) {
		steps = math.Abs(dx)
	} else {
		steps = math.Abs(dy)
	}
	xInc := dx / steps
	yInc := dy / steps
	for s := 0.0; s < steps; s++ {
		p1.X += xInc
		p1.Y += yInc
		v = append(v, NewPoint(p1.X, p1.Y))
	}
	return
}

func AddPoints(p1, p2 *Point) *Point {
	return NewPoint(p1.X+p2.X, p1.Y+p2.Y)
}

func GetEndPoint(p *Point, l float64, a float64) *Point {
	x := p.X + l*math.Cos(a)
	y := p.Y + l*math.Sin(a)
	return NewPoint(x, y)
}
