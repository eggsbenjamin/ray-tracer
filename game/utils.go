package game

import (
	"math"
)

//	get the quadrant of the angle
func GetQuadrant(a float64) int {
	if a < math.Pi/2 {
		return 1
	}
	if a >= math.Pi/2 && a < math.Pi {
		return 2
	}
	if a >= math.Pi && a < 3*(math.Pi/2) {
		return 3
	}
	return 4
}

//	dda algorithm implemented to obtain array of points on a line where
//	the line intersects with a an integer on either the x or y axis
func DDA(p1, p2 Point) (v []*Point) {
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

//	given a starting point, a distance and an angle,
//	calculate the point at the other end of the vector
func GetEndPoint(p *Point, l float64, a float64) *Point {
	x := p.X + l*math.Cos(a)
	y := p.Y + l*math.Sin(a)
	return NewPoint(x, y)
}

//	given a starting point and an angle, calculate the
//  first point along the vector at which x is an integer
func GetFirstXIntersect(p *Point, a float64) *Point {
	var (
		x, y, a2 float64
	)
	q := GetQuadrant(a)
	switch q {
	case 1:
		x = math.Ceil(p.X)
		a2 = a
		y = p.Y + (x-p.X)*math.Tan(a2)
	case 2:
		x = math.Floor(p.X)
		a2 = math.Pi - a
		y = p.Y + (p.X-x)*math.Tan(a2)
	case 3:
		x = math.Floor(p.X)
		a2 = a - math.Pi
		y = p.Y + (x-p.X)*math.Tan(a2)
	case 4:
		x = math.Ceil(p.X)
		a2 = 2*(math.Pi) - a
		y = p.Y + (p.X-x)*math.Tan(a2)
	}
	return NewPoint(x, y)
}

//	given a starting point, an angle and a length, calculate
//	all points along the vector at which x is an integer
func GetXIntersects(p *Point, a, l float64) []*Point {
	ad := float64(1)
	op := math.Tan(a)
	hy := math.Sqrt(math.Pow(ad, 2) + math.Pow(op, 2))
	q := GetQuadrant(a)
	if q == 2 || q == 3 {
		op = -op
		ad = -ad
	}
	s := int(l / hy)
	f := GetFirstXIntersect(p, a)
	is := []*Point{f}
	for i := 0; i < s; i++ {
		x := is[len(is)-1].X + ad
		y := is[len(is)-1].Y + op
		is = append(is, NewPoint(x, y))
	}
	return is
}

//	given a starting point and an angle, calculate the
//  first point along the vector at which y is an integer
func GetFirstYIntersect(p *Point, a float64) *Point {
	var (
		x, y float64
	)
	if a > math.Pi {
		y = float64(int(p.Y) + 1)
	} else {
		y = float64(int(p.Y) - 1)
	}
	x = p.X + (p.Y-y)/math.Tan(a)
	return NewPoint(x, y)
}
