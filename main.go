package main

import (
	"fmt"
	"image/jpeg"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/eggsbenjamin/ray-tracer/game"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

func init() {
	if retVal := img.Init(img.INIT_JPG); retVal == 0 {
		panic("Handle this better! " + img.GetError().Error())
	}
}

const (
	defaultWidth, defaultHeight int = 640, 480
	defaultFPS                  int = 30
)

var (
	width, height int = defaultWidth, defaultHeight
	FPS           int = defaultFPS
)

func init() {
	widthStr, heightStr, FPSStr := os.Getenv("WIDTH"), os.Getenv("HEIGHT"), os.Getenv("FPS")

	widthV, err := strconv.Atoi(widthStr)
	if err != nil {
		fmt.Println("Using default width")
		widthV = defaultWidth
	}
	width = widthV

	heightV, err := strconv.Atoi(heightStr)
	if err != nil {
		fmt.Println("Using default height")
		heightV = defaultHeight
	}
	height = heightV

	FPSV, err := strconv.Atoi(FPSStr)
	if err != nil {
		fmt.Println("Using default FPS")
		FPSV = defaultFPS
	}
	FPS = FPSV

	fmt.Printf("FPS: %d\n", FPS)
	fmt.Printf("Resolution: %dx%d\n", width, height)
}

func clear(r *sdl.Renderer) {
	if err := r.Clear(); err != nil {
		panic(err)
	}
}

func run(l *game.Level, pl *game.Player, c *game.Camera, w *sdl.Window, r *sdl.Renderer) {
	var done bool
	tick := time.NewTicker(time.Second / 30)

	for !done {
		select {
		case <-tick.C:
			clear(r)
			handleEvents(pl, l, &done)
			render(c, l, w, r)
			r.Present()
		}
	}
}

func render(c *game.Camera, l *game.Level, win *sdl.Window, r *sdl.Renderer) {
	w, h := win.GetSize()
	mw, mh := l.GetSize()
	c.Render(w, h, mw, mh, 10.0, r)
}

func handleKeyDownEvent(pl *game.Player, l *game.Level, e *sdl.KeyDownEvent) {
	d := 0.2
	switch e.Keysym.Sym {
	case sdl.K_w:
		n := game.GetEndPoint(pl.Pos(), d, pl.Dir())
		if ok := l.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_s:
		n := game.GetEndPoint(pl.Pos(), -d, pl.Dir())
		if ok := l.Walkable(n); ok {
			pl.Move(n)
		}
	case sdl.K_a:
		pl.Rotate(-0.1)
	case sdl.K_d:
		pl.Rotate(0.1)
	}
}

func handleEvents(pl *game.Player, l *game.Level, done *bool) {
	for {
		e := sdl.PollEvent()
		if e == nil {
			return
		}

		switch e.(type) {
		case *sdl.KeyDownEvent:
			if k, ok := e.(*sdl.KeyDownEvent); ok {
				handleKeyDownEvent(pl, l, k)
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

	sdl.ShowCursor(0) // hide cursor

	pt := game.NewPalette()
	pt[0] = game.BLACK
	pt[1] = game.BLUE
	pt[2] = game.RED

	doorTx, err := img.LoadTexture(r, "./assets/sci-fi_panel_texture-door.jpg")
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	wallTx, err := img.LoadTexture(r, "./assets/sci-fi_floor_texture-wall.jpg")
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	door, err := os.Open("./assets/sci-fi_panel_texture-door.jpg")
	if err != nil {
		panic(err)
	}
	doorImg, err := jpeg.Decode(door)
	if err != nil {
		panic(err)
	}

	wall, err := os.Open("./assets/sci-fi_floor_texture-wall.jpg")
	if err != nil {
		panic(err)
	}

	wallImg, err := jpeg.Decode(wall)
	if err != nil {
		panic(err)
	}

	wallTPt := game.NewTexturePalette()
	wallTPt[1] = game.NewTexture(doorImg, doorTx)
	wallTPt[2] = game.NewTexture(wallImg, wallTx)

	walls := game.NewMap(5, 5, wallTPt)
	walls.Grid = [][]int{
		{2, 2, 2, 2, 1, 2, 2, 2, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 0, 2, 2, 0, 2, 2, 0, 2},
		{2, 0, 2, 2, 0, 2, 2, 0, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 0, 2, 2, 0, 2, 2, 0, 2},
		{2, 0, 2, 2, 0, 2, 2, 0, 2},
		{2, 0, 0, 0, 0, 0, 0, 0, 2},
		{2, 2, 2, 2, 1, 2, 2, 2, 2},
	}

	floorTx, err := img.LoadTexture(r, "./assets/metal_floor_texture-floor.jpg")
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	fl, err := os.Open("./assets/metal_floor_texture-floor.jpg")
	if err != nil {
		panic(err)
	}

	floorImg, err := jpeg.Decode(fl)
	if err != nil {
		panic(err)
	}

	floorTpt := game.NewTexturePalette()
	floorTpt[0] = game.NewTexture(floorImg, floorTx)
	floor := game.NewMap(5, 5, floorTpt)

	ceilTx, err := img.LoadTexture(r, "./assets/rusty_floor_texture-floor2.jpg")
	if err != nil {
		panic("Handle this better! " + err.Error())
	}

	ce, err := os.Open("./assets/rusty_floor_texture-floor2.jpg")
	if err != nil {
		panic(err)
	}

	ceilImg, err := jpeg.Decode(ce)
	if err != nil {
		panic(err)
	}

	ceilTpt := game.NewTexturePalette()
	ceilTpt[0] = game.NewTexture(ceilImg, ceilTx)
	ceil := game.NewMap(5, 5, ceilTpt)

	level := game.NewLevel(game.NewPoint(1.5, 1.5), floor, walls, ceil)
	pl := game.NewPlayer(level.StartingPos, math.Pi/4)
	c := game.NewCamera(pl, level, 0.5, 0.75)

	run(level, pl, c, w, r)
}
