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
			drawMap(m, w, r)
			drawPlayer(pl, w, r)
			r.Present()
		}
	}
}

func drawMap(m *game.Map, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	b := &sdl.Rect{0, 0, int32(w), int32(h)}
	r.SetDrawColor(0, 92, 161, 1)
	r.FillRect(b)

	r.SetDrawColor(255, 255, 0, 1)
	for i := 1; i < m.Width; i++ {
		r.DrawLine(i*w/m.Width, h, i*w/m.Width, 0)
	}
	for i := 1; i < m.Height; i++ {
		r.DrawLine(0, i*h/m.Height, w, i*h/m.Height)
	}
}

func drawPlayer(pl *game.Player, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	var xSc float64 = float64(w) / 5
	var ySc float64 = float64(h) / 5
	l := game.GetEndPoint(pl.Pos, pl.Cam.FocalLength, pl.Dir-(math.Pi/4))
	rg := game.GetEndPoint(pl.Pos, pl.Cam.FocalLength, pl.Dir+(math.Pi/4))
	r.SetDrawColor(255, 255, 255, 1)
	r.DrawPoint(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc))
	r.DrawLine(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc), int(l.X*xSc), int(l.Y*ySc))
	r.DrawLine(int(pl.Pos.X*xSc), int(pl.Pos.Y*ySc), int(rg.X*xSc), int(rg.Y*ySc))
}

func handleKeyUpEvent(pl *game.Player, e *sdl.KeyUpEvent) {
}

func handleKeyDownEvent(pl *game.Player, e *sdl.KeyDownEvent) {
	switch e.Keysym.Sym {
	case sdl.K_UP:
		pl.Move(0.2)
	case sdl.K_DOWN:
		pl.Move(-0.2)
	case sdl.K_LEFT:
		pl.Rotate(-0.2)
	case sdl.K_RIGHT:
		pl.Rotate(0.2)
	}
}

func handleEvents(pl *game.Player, events chan sdl.Event, done chan<- bool) {
	for {
		e := <-events
		switch e.(type) {
		case *sdl.KeyUpEvent:
			if k, ok := e.(*sdl.KeyUpEvent); ok {
				handleKeyUpEvent(pl, k)
			}
		case *sdl.KeyDownEvent:
			if k, ok := e.(*sdl.KeyDownEvent); ok {
				handleKeyDownEvent(pl, k)
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
	m := game.NewMap(5, 5)

	events := make(chan sdl.Event)
	done := make(chan bool)
	go handleEvents(pl, events, done)
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
