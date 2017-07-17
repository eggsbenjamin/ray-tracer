package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Camera struct {
	Pos         *Point
	Angle       float64
	FocalLength float64
	FOV         float64
}

//	constructor
func NewCamera(pos *Point, a, fl, di float64) *Camera {
	fov := 2 * math.Atan((di/2)/fl)
	return &Camera{
		Pos:         pos,
		Angle:       a,
		FocalLength: fl,
		FOV:         fov,
	}
}

func (c *Camera) Render(w, h, mw, mh int, l float64, r *sdl.Renderer) {
	var xSc float64 = float64(w) / float64(mw)
	var ySc float64 = float64(h) / float64(mh)
	d := c.FOV / float64(w)
	r.SetDrawColor(255, 255, 255, 1)
	for i := 0; i < w; i++ {
		p := GetEndPoint(c.Pos, l, (c.Angle-math.Pi/4)+(float64(i)*d))
		v := NewVector(*c.Pos, *p)
		a := v.Start()
		b := v.End()
		r.DrawLine(int(a.X*xSc), int(a.Y*ySc), int(b.X*xSc), int(b.Y*ySc))
		for s := range *v {
			//	TODO - detect if the ray has hit a wall...
		}
	}
}
