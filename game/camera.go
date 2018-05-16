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
		ray := NewRay(c.Focus.Pos(), angle, l)
		p, v, _, xOffset := c.getNearestHit(ray)
		if p != nil {
			ra := (c.FOV / 2) - float64(x)*d
			dp := DistanceBetweenPoints(c.Focus.Pos(), p)
			di := dp * math.Cos(ra)
			sh := 512 / di
			y1, y2 := (h/2)-int(sh/2), (h/2)+int(sh/2)
			tex := c.Level.Walls.TexturePalette[v]
			_, _, iw, ih, _ := tex.Query()
			texX := float64(iw) * xOffset

			if err := renderer.Copy(tex.Texture, &sdl.Rect{int32(texX), 0, 1, ih}, &sdl.Rect{int32(x), int32(y1), 1, int32(y2 - y1)}); err != nil {
				panic(err)
			}

			/*
				ppl := (math.Tan(c.FOV/2) * c.FocalLength) * 2
				cf := ppl / float64(w)
				dppp := math.Sqrt(math.Pow(math.Abs(float64(w/2-x))*cf, 2) + math.Pow(c.FocalLength, 2))

				for i := y2; i < h; i++ {
					dcpp := float64(i-(h/2)) * cf
					ta := dppp / dcpp
					df := (float64(512) * cf) * ta

					fp := NewPoint(c.Focus.Pos().X+df*math.Cos(angle), c.Focus.Pos().Y+df*math.Sin(angle))

					floorTex := c.Level.Floor.GetTextureAtPoint(NewPoint(1, 1))
					texX := (fp.X - math.Floor(fp.X)) * float64(floorTex.Image.Bounds().Dx())
					texY := (fp.Y - math.Floor(fp.Y)) * float64(floorTex.Image.Bounds().Dy())

					r32, g32, b32, a32 := floorTex.At(int(texX), int(texY)).RGBA()
					r, g, b, a := uint8(r32>>8), uint8(g32>>8), uint8(b32>>8), uint8(a32>>8)

					renderer.SetDrawColor(r, g, b, a)
					renderer.DrawPoint(x, i) // draw floor

					ceilTex := c.Level.Ceiling.GetTextureAtPoint(NewPoint(1, 1))
					texX = (fp.X - math.Floor(fp.X)) * float64(ceilTex.Image.Bounds().Dx())
					texY = (fp.Y - math.Floor(fp.Y)) * float64(ceilTex.Image.Bounds().Dy())

					r32, g32, b32, a32 = ceilTex.At(int(texX), int(texY)).RGBA()
					r, g, b, a = uint8(r32>>8), uint8(g32>>8), uint8(b32>>8), uint8(a32>>8)

					renderer.SetDrawColor(r, g, b, a)
					renderer.DrawPoint(x, y1-(i-y2)) // draw ceiling
				}
			*/
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
	col := BLACK
	r.SetDrawColor(col.R, col.G, col.B, col.A)
	r.FillRect(&sdl.Rect{
		0, 0, int32(w), int32(h / 2),
	})
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
			xv = c.Level.Walls.GetValueAtPoint(NewPoint(xi.X-1, xi.Y))
		} else {
			xv = c.Level.Walls.GetValueAtPoint(xi)
		}
		if xv > 0 {
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
	xd := DistanceBetweenPoints(c.Focus.Pos(), x)
	yd := DistanceBetweenPoints(c.Focus.Pos(), y)
	if xd < yd {
		return x, xv, false, x.Y - math.Floor(x.Y)
	}
	return y, yv, true, y.X - math.Floor(y.X)
}
