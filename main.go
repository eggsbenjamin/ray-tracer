package main

import (
	"math"
	"os"
	"runtime"
	"time"

	"github.com/eggsbenjamin/ray-tracer/game"
	"github.com/veandco/go-sdl2/sdl"
)

func clear(r *sdl.Renderer) {
	wr := &sdl.Rect{0, 0, 500, 500}
	r.SetDrawColor(0, 0, 0, 0)
	r.FillRect(wr)
}

func run(m *game.Map, pl *game.Player, w *sdl.Window, r *sdl.Renderer) {
	tick := time.NewTicker(time.Second / 30)
	for {
		select {
		case <-tick.C:
			clear(r)
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
	w, h := win.GetSize()
	mw, mh := m.GetSize()
	pl.Cam.Render(w, h, mw, mh, 10.0, r)
}

func drawCeiling(win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	rc := &sdl.Rect{0, 0, int32(w), int32(h / 2)}
	col := game.GREY
	r.SetDrawColor(col.R, col.G, col.B, col.A)
	r.FillRect(rc)
}

func handleKeyDownEvent(pl *game.Player, m *game.Map, e *sdl.KeyDownEvent) {
	d := 0.05
	switch e.Keysym.Sym {
	case sdl.K_UP:
		n := game.GetEndPoint(pl.Pos, d, pl.Dir)
		if ok := m.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_DOWN:
		n := game.GetEndPoint(pl.Pos, -d, pl.Dir)
		if ok := m.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_LEFT:
		pl.Rotate(-0.1)
	case sdl.K_RIGHT:
		pl.Rotate(0.1)
	}
}

func handleEvents(pl *game.Player, m *game.Map, events chan sdl.Event, done chan<- bool) {
	for {
		e := <-events
		switch e.(type) {
		case *sdl.KeyDownEvent:
			if k, ok := e.(*sdl.KeyDownEvent); ok {
				handleKeyDownEvent(pl, m, k)
			}
		case *sdl.QuitEvent:
			done <- true
		}
	}
}

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
	w, r, err := sdl.CreateWindowAndRenderer(500, 500, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer w.Destroy()

	pt := game.NewPalette()
	pt[0] = game.BLACK
	pt[1] = game.BLUE
	pt[2] = game.RED
	m := game.NewMap(5, 5, pt)
	m.Grid = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 2, 0, 2, 2, 2, 0, 0, 1},
		{1, 0, 2, 0, 0, 0, 2, 0, 0, 1},
		{1, 0, 2, 0, 0, 0, 2, 0, 0, 1},
		{1, 0, 2, 2, 0, 2, 2, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}
	pl := game.NewPlayer(m, 2.5, 2.5, math.Pi/4)

	events := make(chan sdl.Event)
	done := make(chan bool)
	go handleEvents(pl, m, events, done)
	go run(m, pl, w, r)

	runtime.LockOSThread()
	for {
		select {
		case events <- sdl.WaitEvent():
		case <-done:
			os.Exit(0)
		}
	}
}
