package main

import (
	"image/jpeg"
	"math"
	"os"
	"time"

	"github.com/eggsbenjamin/ray-tracer/game"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	width, height int = 640, 480
)

func clear(r *sdl.Renderer) {
	wr := &sdl.Rect{0, 0, 500, 500}
	r.SetDrawColor(0, 0, 0, 0)
	r.FillRect(wr)
}

func run(m *game.Map, pl *game.Player, w *sdl.Window, r *sdl.Renderer) {
	var done bool
	tick := time.NewTicker(time.Second / 30)

	for !done {
		select {
		case <-tick.C:
			clear(r)
			handleEvents(pl, m, &done)
			drawCeiling(w, r)
			drawPlayer(pl, m, w, r)
			drawMap(m, pl, w, r)
			r.Present()
		}
	}
}

func drawMap(m *game.Map, pl *game.Player, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	mw, mh := m.GetSize()
	tw := (w / mw) / 4
	th := (h / mh) / 4
	for y := 0; y < mh; y++ {
		for x := 0; x < mw; x++ {
			col := m.Palette[m.Grid[x][y]]
			rc := &sdl.Rect{int32(x * tw), int32(y * th), int32(tw), int32(th)}
			r.SetDrawColor(col.R, col.G, col.B, col.A)
			r.FillRect(rc)
		}
	}
	xSc := float64(w) / float64(mw)
	ySc := float64(h) / float64(mh)
	d := pl.Cam.FOV / float64(w)
	r.SetDrawColor(255, 255, 255, 1)
	for i := 0; i <= w; i++ {
		p := game.GetEndPoint(pl.Pos, 2.5, (pl.Dir-math.Pi/4)+(float64(i)*d))
		r.DrawLine(int(xSc*pl.Pos.X)/4, int(ySc*pl.Pos.Y)/4, int(xSc*p.X)/4, int(ySc*p.Y)/4)
	}
}

func drawPlayer(pl *game.Player, m *game.Map, win *sdl.Window, r *sdl.Renderer) {
	mw, mh := m.GetSize()
	pl.Cam.Render(width, height, mw, mh, 10.0, r)
}

func drawCeiling(win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	rc := &sdl.Rect{0, 0, int32(w), int32(h / 2)}
	col := game.GREY
	r.SetDrawColor(col.R, col.G, col.B, col.A)
	r.FillRect(rc)
}

func handleKeyDownEvent(pl *game.Player, m *game.Map, e *sdl.KeyDownEvent) {
	d := 0.2
	switch e.Keysym.Sym {
	case sdl.K_w:
		n := game.GetEndPoint(pl.Pos, d, pl.Dir)
		if ok := m.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_s:
		n := game.GetEndPoint(pl.Pos, -d, pl.Dir)
		if ok := m.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_a:
		pl.Rotate(-0.1)
	case sdl.K_d:
		pl.Rotate(0.1)
	}
}

func handleEvents(pl *game.Player, m *game.Map, done *bool) {
	for {
		e := sdl.PollEvent()
		if e == nil {
			return
		}

		switch e.(type) {
		case *sdl.KeyDownEvent:
			if k, ok := e.(*sdl.KeyDownEvent); ok {
				handleKeyDownEvent(pl, m, k)
			}
		case *sdl.QuitEvent:
			*done = true
		}
	}
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	w, r, err := sdl.CreateWindowAndRenderer(width, height, sdl.WINDOW_FULLSCREEN_DESKTOP)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()
	defer r.Destroy()

	displayMode := &sdl.DisplayMode{}
	if err := sdl.GetDesktopDisplayMode(0, displayMode); err != nil {
		panic(err)
	}

	if err := r.SetScale(float32(displayMode.W)/float32(width), float32(displayMode.H)/float32(height)); err != nil {
		panic(err)
	}

	pt := game.NewPalette()
	pt[0] = game.BLACK
	pt[1] = game.BLUE
	pt[2] = game.RED

	stoneWall, err := os.Open(game.STONE_WALL_PATH)
	if err != nil {
		panic(err)
	}
	stoneWallTexture, err := jpeg.Decode(stoneWall)
	if err != nil {
		panic(err)
	}

	brickWall, err := os.Open(game.ORANGE_STONE_WALL)
	if err != nil {
		panic(err)
	}
	brickWallTexture, err := jpeg.Decode(brickWall)
	if err != nil {
		panic(err)
	}

	tPt := game.NewTexturePalette()
	tPt[2] = brickWallTexture
	tPt[1] = stoneWallTexture

	m := game.NewMap(5, 5, pt, tPt, nil)
	m.Grid = [][]int{
		{2, 2, 2, 2, 2, 2, 2, 2},
		{2, 0, 0, 0, 0, 0, 0, 2},
		{2, 0, 2, 0, 2, 2, 0, 2},
		{2, 0, 2, 0, 0, 2, 0, 2},
		{2, 0, 2, 0, 0, 2, 0, 2},
		{2, 0, 2, 2, 0, 2, 0, 2},
		{2, 0, 0, 0, 0, 0, 0, 2},
		{2, 2, 2, 2, 2, 2, 2, 2},
	}
	pl := game.NewPlayer(m, 1.5, 1.5, math.Pi/4)

	run(m, pl, w, r)
}
