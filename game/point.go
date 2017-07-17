package game

type Point struct {
	X float64
	Y float64
}

//	constructor
func NewPoint(x, y float64) *Point {
	return &Point{
		X: x,
		Y: y,
	}
}
