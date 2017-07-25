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

//	adds two point together to produce a new point
func AddPoints(p1, p2 *Point) *Point {
	return NewPoint(p1.X+p2.X, p1.Y+p2.Y)
}

//	calculates the distance between two points
func DistanceBetweenPoints(p1, p2 *Point) float64 {
	return math.Sqrt(math.Pow(p2.X-p1.X, 2) + math.Pow(p2.Y-p1.Y, 2))
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

//	given a starting point and an angle, calculate the
//  first point along the vector at which y is an integer
func GetFirstYIntersect(p *Point, a float64) *Point {
	var (
		x, y, a2 float64
		q        = GetQuadrant(a)
	)
	switch q {
	case 1:
		y = math.Ceil(p.Y)
		a2 = a
		x = p.X + (y-p.Y)/math.Tan(a2)
	case 2:
		y = math.Ceil(p.Y)
		a2 = math.Pi - a
		x = p.X + (p.Y-y)/math.Tan(a2)
	case 3:
		y = math.Floor(p.Y)
		a2 = a - math.Pi
		x = p.X + (y-p.Y)/math.Tan(a2)
	case 4:
		y = math.Floor(p.Y)
		a2 = 2*(math.Pi) - a
		x = p.X + (p.Y-y)/math.Tan(a2)
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

//	given a starting point, an angle and a length, calculate
//	all points along the vector at which y is an integer
func GetYIntersects(p *Point, a, l float64) []*Point {
	op := float64(1)
	ad := op / math.Tan(a)
	hy := math.Sqrt(math.Pow(ad, 2) + math.Pow(op, 2))
	q := GetQuadrant(a)
	if q == 3 || q == 4 {
		op = -op
		ad = -ad
	}
	s := int(l / hy)
	f := GetFirstYIntersect(p, a)
	is := []*Point{f}
	for i := 0; i < s; i++ {
		x := is[len(is)-1].X + ad
		y := is[len(is)-1].Y + op
		is = append(is, NewPoint(x, y))
	}
	return is
}
