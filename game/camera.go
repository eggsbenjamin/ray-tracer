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
	for x := 1; x <= w; x++ {
		a := NormaliseAngle((c.Angle - c.FOV/2) + (float64(x) * d))
		ry := NewRay(c.Pos, a, l)
		p, v, _ := c.getNearestHit(ry)
		if p != nil {
			di := DistanceBetweenPoints(c.Pos, p)
			sh := 128 / di * c.FocalLength
			col := c.Map.Palette[v]
			r.SetDrawColor(col.R, col.G, col.B, col.A)
			r.DrawLine(x, (h/2)-int(sh/2), x, (h/2)+int(sh/2))
		}
	}
}

func (c *Camera) getNearestHit(r *Ray) (*Point, int, bool) {
	var (
		xv, yv int
		x, y   *Point = nil, nil
		q             = GetQuadrant(r.angle)
	)
	xis := r.XIntersects()
	for _, xi := range xis {
		if q == 2 || q == 3 {
			xv = c.Map.GetValueAtPoint(NewPoint(xi.X-1, xi.Y))
		} else {
			xv = c.Map.GetValueAtPoint(xi)
		}
		if xv > 0 {
			x = xi
			break
		}
	}
	yis := r.YIntersects()
	for _, yi := range yis {
		if q > 2 {
			yv = c.Map.GetValueAtPoint(NewPoint(yi.X, yi.Y-1))
		} else {
			yv = c.Map.GetValueAtPoint(yi)
		}
		if yv > 0 {
			y = yi
			break
		}
	}
	if x == nil && y == nil {
		return nil, 0, false
	}
	if x == nil {
		return y, yv, true
	}
	if y == nil {
		return x, xv, false
	}
	xd := DistanceBetweenPoints(c.Pos, x)
	yd := DistanceBetweenPoints(c.Pos, y)
	if xd < yd {
		return x, xv, false
	}
	return y, yv, true
}
