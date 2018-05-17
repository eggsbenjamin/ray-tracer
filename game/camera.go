package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Camera struct {
	Focus       Focusable
	FocalLength float64
	FOV         float64
	Level       *Level
}

//	constructor
func NewCamera(f Focusable, l *Level, fl, di float64) *Camera {
	fov := 2 * math.Atan((di/2)/fl)
	return &Camera{
		Focus:       f,
		FocalLength: fl,
		FOV:         fov,
		Level:       l,
	}
}

func (c *Camera) Render(w, h, mw, mh int, l float64, renderer *sdl.Renderer) {
	c.drawFloor(w, h, GREY, renderer)
	c.drawSky(w, h, renderer)

	d := c.FOV / float64(w)
	for x := 1; x <= w; x++ {
		angle := NormaliseAngle((c.Focus.Dir() - c.FOV/2) + (float64(x) * d))
		ray := NewRay(c.Focus.Pos(), angle, l*2)
		p, v, _, xOffset := c.getNearestHit(ray)
		if p != nil {
			ra := (c.FOV / 2) - float64(x)*d
			dp := DistanceBetweenPoints(c.Focus.Pos(), p)
			di := dp * math.Cos(ra)
			sh := (float64(w) * 0.8) / di
			y1, y2 := (h/2)-int(sh/2), (h/2)+int(sh/2)
			tex := v.Texture()
			_, _, iw, ih, _ := v.Texture().Query()
			texX := float64(iw) * xOffset

			if err := renderer.Copy(tex.Texture, &sdl.Rect{int32(texX), 0, 1, ih}, &sdl.Rect{int32(x), int32(y1), 1, int32(y2 - y1)}); err != nil {
				panic(err)
			}
		}
	}
}

func (c *Camera) drawFloor(w, h int, col *Colour, r *sdl.Renderer) {
	r.SetDrawColor(col.R, col.G, col.B, col.A)
	r.FillRect(&sdl.Rect{
		0, int32(h / 2), int32(w), int32(h / 2),
	})
}

func (c *Camera) drawSky(w, h int, r *sdl.Renderer) {
	col := BLACK // TODO - skybox??
	r.SetDrawColor(col.R, col.G, col.B, col.A)
	r.FillRect(&sdl.Rect{
		0, 0, int32(w), int32(h / 2),
	})
}

func (c *Camera) getNearestHit(r *Ray) (*Point, Tile, bool, float64) {
	var (
		xv, yv Tile
		x, y   *Point = nil, nil
		q             = GetQuadrant(r.angle)
	)
	xis := r.XIntersects()
	for _, xi := range xis {
		if q == 2 || q == 3 {
			xv = c.Level.Walls.GetValueAtPoint(NewPoint(xi.X-1, xi.Y))
		} else {
			xv = c.Level.Walls.GetValueAtPoint(xi)
		}
		if xv != nil && xv.Visable() {
			x = xi
			break
		}
	}
	yis := r.YIntersects()
	for _, yi := range yis {
		if q > 2 {
			yv = c.Level.Walls.GetValueAtPoint(NewPoint(yi.X, yi.Y-1))
		} else {
			yv = c.Level.Walls.GetValueAtPoint(yi)
		}
		if yv != nil && yv.Visable() {
			y = yi
			break
		}
	}
	if x == nil && y == nil {
		return nil, nil, false, 0
	}
	if x == nil {
		return y, yv, true, y.X - math.Floor(y.X)
	}
	if y == nil {
		return x, xv, false, x.Y - math.Floor(x.Y)
	}
	xd := DistanceBetweenPoints(c.Focus.Pos(), x)
	yd := DistanceBetweenPoints(c.Focus.Pos(), y)
	if xd < yd {
		return x, xv, false, x.Y - math.Floor(x.Y)
	}
	return y, yv, true, y.X - math.Floor(y.X)
}
