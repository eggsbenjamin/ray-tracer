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
	xSc := float64(w) / float64(mw)
	ySc := float64(h) / float64(mh)
	p := GetEndPoint(c.Pos, l, c.Angle)
	r.SetDrawColor(255, 255, 255, 1)
	r.DrawLine(int(c.Pos.X*xSc), int(c.Pos.Y*ySc), int(p.X*xSc), int(p.Y*ySc))
	/*d := c.FOV / float64(w)
	r.SetDrawColor(255, 255, 255, 1)
	for i := 0; i < w; i++ {
		p := GetEndPoint(c.Pos, l, (c.Angle-math.Pi/4)+(float64(i)*d))
		is := DDA(*c.Pos, *p)
		for _, _ = range is {
		}
		r.DrawLine(int(c.Pos.X*xSc), int(c.Pos.Y*ySc), int(p.X*xSc), int(p.Y*ySc))
	}*/
}
