package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/eggsbenjamin/ray-tracer/asset"
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
	defaultWidth, defaultHeight                  int    = 640, 480
	defaultFPS                                   int    = 30
	fontSize                                     int    = 1000
	soundsPath, musicPath, fontsPath, imagesPath string = "./assets/audio/sounds/", "./assets/audio/music/", "./assets/fonts/", "./assets/images/"
)

var (
	width, height int = defaultWidth, defaultHeight
	FPS           int = defaultFPS
	rTex          *sdl.Texture
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

func run(l *game.Level, pl *game.Player, c *game.Camera, w *sdl.Window, r *sdl.Renderer) {
	var done bool
	tick := time.NewTicker(time.Second / 30)

	for !done {
		select {
		case <-tick.C:
			handleEvents(pl, l, &done)
			update(pl, l)
			render(c, l, w, r)
		}
	}
}

func update(pl *game.Player, l *game.Level) {
	pl.Update(l)
}

func render(c *game.Camera, l *game.Level, win *sdl.Window, r *sdl.Renderer) {
	if err := r.Clear(); err != nil {
		panic(err)
	}

	w, h := win.GetSize()
	mw, mh := l.GetSize()
	c.Render(w, h, mw, mh, 10.0, r)
	if err := r.SetRenderTarget(nil); err != nil {
		panic(err)
	}

	if err := r.Copy(rTex, nil, nil); err != nil {
		panic(err)
	}

	r.Present()
	if err := r.SetRenderTarget(rTex); err != nil {
		panic(err)
	}
}

func handleKeyDownEvent(pl *game.Player, l *game.Level, e *sdl.KeyDownEvent) {
	d := 0.2
	switch e.Keysym.Sym {
	case sdl.K_w:
		pl.SetState("walking-forward")
	case sdl.K_s:
		pl.SetState("walking-backward")
	case 46:
		pl.SetState("strafing-left")
	case 47:
		pl.SetState("strafing-right")
	case sdl.K_a:
		pl.SetState("rotating-left")
	case sdl.K_d:
		pl.SetState("rotating-right")
	case sdl.K_e:
		n := game.GetEndPoint(pl.Pos(), d*3, pl.Dir())
		l.Interact(n, pl)
	}
}

func handleKeyUpEvent(pl *game.Player, l *game.Level, e *sdl.KeyUpEvent) {
	switch e.Keysym.Sym {
	case sdl.K_w:
		pl.UnsetState("walking-forward")
	case sdl.K_s:
		pl.UnsetState("walking-backward")
	case sdl.K_a:
		pl.UnsetState("rotating-left")
	case sdl.K_d:
		pl.UnsetState("rotating-right")
	case 46:
		pl.UnsetState("strafing-left")
	case 47:
		pl.UnsetState("strafing-right")
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
		case *sdl.KeyUpEvent:
			if k, ok := e.(*sdl.KeyUpEvent); ok {
				handleKeyUpEvent(pl, l, k)
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
	tx, err := r.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_TARGET, width, height)
	if err != nil {
		panic(err)
	}
	rTex = tx

	if err := r.SetRenderTarget(rTex); err != nil {
		panic(err)
	}

	assetmgr := asset.NewManager(r, fontsPath, imagesPath, fontSize)
	assetmgr.LoadContent()

	level := initLevel(assetmgr)
	pl := game.NewPlayer(level.StartingPos, math.Pi/4, 0.25, 0.1, nil)
	c := game.NewCamera(pl, level, 0.5, 0.75)

	run(level, pl, c, w, r)
}

func initLevel(assetmgr asset.Manager) *game.Level {
	wl := game.NewWall(assetmgr.GetImage("sci-fi_floor_texture-wall"))
	sWl := game.NewWall(assetmgr.GetImage("metal_floor_texture-floor"))

	walls := game.NewMap(5, 5, nil)
	walls.Grid = [][]game.Tile{
		{wl, wl, wl, wl, game.NewDoor(assetmgr.GetImage("sci-fi_panel_texture-door")), wl, wl, wl, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, wl, wl, wl, game.NewDoor(assetmgr.GetImage("sci-fi_panel_texture-door")), wl, wl, wl, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, sWl, nil, nil, nil, sWl, nil, wl},
		{wl, nil, nil, nil, nil, nil, nil, nil, wl},
		{wl, wl, wl, wl, game.NewDoor(assetmgr.GetImage("sci-fi_panel_texture-door")), wl, wl, wl, wl},
		{wl, wl, wl, nil, nil, sWl, sWl, sWl, wl},
		{wl, wl, wl, nil, nil, sWl, sWl, sWl, wl},
		{wl, wl, wl, nil, nil, nil, nil, nil, wl},
		{wl, wl, wl, nil, nil, wl, wl, nil, wl},
		{wl, wl, wl, nil, nil, wl, wl, nil, wl},
		{wl, wl, wl, nil, nil, nil, nil, nil, wl},
		{wl, wl, wl, wl, wl, wl, wl, wl, wl},
	}

	return game.NewLevel(game.NewPoint(1.5, 1.5), nil, walls, nil, nil)
}

func initWeapons(assetmgr asset.Manager) []game.Weapon {
	hgImg := assetmgr.GetImage("handgun_sprite")
	hgSs := game.NewSpriteSheet(hgImg)
	hgAnim := game.NewAnimation(hgSs, []*sdl.Rect{
		{10, 5, 85, 105},
		{120, 5, 95, 105},
		{240, 5, 85, 105},
		{345, 5, 85, 105},
	})
	_ = hgAnim

	return nil
}
