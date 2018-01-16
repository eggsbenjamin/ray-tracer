package game

import (
	"image"
)

var (
	RED   = NewColour(255, 0, 0, 1)
	BLUE  = NewColour(0, 92, 161, 1)
	BLACK = NewColour(0, 0, 0, 1)
	GREY  = NewColour(149, 144, 182, 1)
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

type WallTexturePalette map[int]image.Image

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
