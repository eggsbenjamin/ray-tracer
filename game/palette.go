package game

import (
	"image"

	"github.com/veandco/go-sdl2/sdl"
)

var (
	RED   = NewColour(255, 0, 0, 1)
	BLUE  = NewColour(0, 92, 161, 1)
	BLACK = NewColour(0, 0, 0, 1)
	GREY  = NewColour(128, 128, 128, 1)
)

type Palette map[int]*Colour

//	constructor
func NewPalette() Palette {
	return Palette{}
}

type Colour struct {
	R, G, B, A uint8
}

//	constructor
func NewColour(r, g, b, a uint8) *Colour {
	return &Colour{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

const (
	ORANGE_STONE_WALL = "./assets/orange_stone_wall.jpg"
	STONE_WALL_PATH   = "./assets/stone_wall.jpg"
	GRASS_PATH        = "./assets/grass.jpg"
)

type Texture struct {
	image.Image
	*sdl.Texture
}

func NewTexture(img image.Image, tx *sdl.Texture) *Texture {
	return &Texture{
		img,
		tx,
	}
}

type WallTexturePalette map[int]*Texture

// constructor
func NewTexturePalette() WallTexturePalette {
	return WallTexturePalette{}
}

type WalkableTexture struct {
	Floor   image.Image
	Ceiling image.Image
}

// constructor
func NewWalkableTexture(f image.Image, c image.Image) *WalkableTexture {
	return &WalkableTexture{
		Floor:   f,
		Ceiling: c,
	}
}

type WalkableTexturePalette map[int]*WalkableTexture

// constructor
func NewWalkableTexturePalette() WalkableTexturePalette {
	return WalkableTexturePalette{}
}
