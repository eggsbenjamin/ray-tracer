package game

import (
	"math"
)

type Camera struct {
	Pos         *Point
	Angle       float64
	FocalLength float64
	FOV         float64
}

//	constructor
func NewCamera(pos *Point, a, fl, di float64) *Camera {
	fov := 2 * ((math.Atan((di / 2) / fl)) * 180) / math.Pi
	return &Camera{
		Pos:         pos,
		Angle:       a,
		FocalLength: fl,
		FOV:         fov,
	}
}
