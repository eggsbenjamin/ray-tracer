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
	Map         *Map
}

//	constructor
func NewCamera(pos *Point, m *Map, a, fl, di float64) *Camera {
	fov := 2 * math.Atan((di/2)/fl)
	return &Camera{
		Pos:         pos,
		Angle:       a,
		FocalLength: fl,
		FOV:         fov,
		Map:         m,
	}
}

func (c *Camera) Render(w, h, mw, mh int, l float64, r *sdl.Renderer) {
	d := c.FOV / float64(w)
	r.SetDrawColor(255, 255, 255, 1)
	for x := 0; x <= w; x++ {
		a := (c.Angle - c.FOV/2) + (float64(x) * d)
		ry := NewRay(c.Pos, a, l)
		p, _ := c.getNearestHit(ry)
		if p != nil {
			di := DistanceBetweenPoints(c.Pos, p)
			sh := 128 / di * c.FocalLength
			r.DrawLine(x, (h/2)-int(sh/2), x, (h/2)+int(sh/2))
		}
	}
}

func (c *Camera) getNearestHit(r *Ray) (*Point, bool) {
	var (
		x, y *Point = nil, nil
	)
	xis := r.XIntersects()
	for _, xi := range xis {
		v := c.Map.GetValueAtPoint(xi)
		if v > 0 {
			x = xi
			break
		}
	}
	yis := r.YIntersects()
	for _, yi := range yis {
		v := c.Map.GetValueAtPoint(yi)
		if v > 0 {
			y = yi
			break
		}
	}
	if x == nil && y == nil {
		return nil, false
	}
	if x == nil {
		return y, true
	}
	if y == nil {
		return x, false
	}
	xd := DistanceBetweenPoints(c.Pos, x)
	yd := DistanceBetweenPoints(c.Pos, y)
	if xd < yd {
		return x, false
	} else {
		return y, true
	}
}
