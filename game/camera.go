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

func (c *Camera) Render(w, h, mw, mh int, l float64, renderer *sdl.Renderer) {
	d := c.FOV / float64(w)
	for x := 1; x <= w; x++ {
		angle := NormaliseAngle((c.Angle - c.FOV/2) + (float64(x) * d))
		ray := NewRay(c.Pos, angle, l)
		p, v, yHit, xOffset := c.getNearestHit(ray)
		if p != nil {
			ra := (c.FOV / 2) - float64(x)*d
			di := DistanceBetweenPoints(c.Pos, p) * math.Cos(ra)
			sh := 128 / di
			y1, y2 := (h/2)-int(sh/2), (h/2)+int(sh/2)
			tex := c.Map.TexturePalette[v]

			for i := y1; i < y2; i++ {
				texX := float64(tex.Bounds().Dx()) * xOffset
				texY := float64((i - y1)) * float64(tex.Bounds().Dy()) / sh
				r32, g32, b32, a32 := tex.At(int(texX), int(texY)).RGBA()
				r, g, b, a := uint8(r32>>8), uint8(g32>>8), uint8(b32>>8), uint8(a32>>8)
				if yHit {
					r, g, b, a = (r >> 1), (g >> 1), (b >> 1), (a >> 1)
				}
				renderer.SetDrawColor(r, g, b, a)
				renderer.DrawPoint(x, i)
			}
		}
	}
}

func (c *Camera) getNearestHit(r *Ray) (*Point, int, bool, float64) {
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
		return nil, 0, false, 0
	}
	if x == nil {
		return y, yv, true, y.X - math.Floor(y.X)
	}
	if y == nil {
		return x, xv, false, x.Y - math.Floor(x.Y)
	}
	xd := DistanceBetweenPoints(c.Pos, x)
	yd := DistanceBetweenPoints(c.Pos, y)
	if xd < yd {
		return x, xv, false, x.Y - math.Floor(x.Y)
	}
	return y, yv, true, y.X - math.Floor(y.X)
}
