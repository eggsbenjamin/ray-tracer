package game

import (
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

type WallTexturePalette map[int]*sdl.Texture

// constructor
func NewTexturePalette() WallTexturePalette {
	return WallTexturePalette{}
}
