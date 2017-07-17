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
			drawPlayer(pl, m, w, r)
			r.Present()
		}
	}
}

func drawMap(m *game.Map, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	mw, mh := m.GetSize()
	tw := w / mw
	th := h / mh
	for y := 0; y < mh; y++ {
		for x := 0; x < mw; x++ {
			col := m.Palette[m.Grid[x][y]]
			rc := &sdl.Rect{int32(x * tw), int32(y * th), int32(tw), int32(th)}
			r.SetDrawColor(col.R, col.G, col.B, col.A)
			r.DrawRect(rc)
		}
	}
}

func drawPlayer(pl *game.Player, m *game.Map, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	mw, mh := m.GetSize()
	pl.Cam.Render(w, h, mw, mh, 10.0, r)
	/*
		var xSc float64 = float64(w) / float64(mw)
		var ySc float64 = float64(h) / float64(mh)
		l := game.GetEndPoint(pl.Pos, pl.Cam.FocalLength, pl.Dir-(math.Pi/4))
		rg := game.GetEndPoint(pl.Pos, pl.Cam.FocalLength, pl.Dir+(math.Pi/4))
		r.SetDrawColor(255, 255, 255, 1)
		r.DrawPoint(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc))
		r.DrawLine(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc), int(l.X*xSc), int(l.Y*ySc))
		r.DrawLine(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc), int(rg.X*xSc), int(rg.Y*ySc))
	*/
}

func handleKeyDownEvent(pl *game.Player, m *game.Map, e *sdl.KeyDownEvent) {
	d := 0.2
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
		pl.Rotate(-0.2)
	case sdl.K_RIGHT:
		pl.Rotate(0.2)
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

	pl := game.NewPlayer(2.5, 2.5, math.Pi/4)
	pt := game.NewPalette()
	pt[0] = game.BLACK
	pt[1] = game.BLUE
	m := game.NewMap(5, 5, pt)
	m.Grid = [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
	}

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
