package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type Map struct {
	Grid           [][]Tile
	TexturePalette WallTexturePalette
}

//	constructor
func NewMap(w, h int, tp WallTexturePalette) *Map {
	grid := make([][]Tile, w)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]Tile, h)
	}
	return &Map{
		Grid:           grid,
		TexturePalette: tp,
	}
}

//	returns the width and height
func (m *Map) GetSize() (int, int) {
	return m.width(), m.height()
}

//	gets the value of a square containing a given point
func (m *Map) GetValueAtPoint(p *Point) Tile {
	x := int(math.Floor(p.X))
	y := int(math.Floor(p.Y))
	w, h := m.GetSize()
	if x < 0 || y >= w || y < 0 || x >= h { // out of bounds...
		return nil
	}
	return m.Grid[x][y]
}

func (m *Map) GetTextureAtPoint(p *Point) *sdl.Texture {
	if v := m.GetValueAtPoint(p); v != nil {
		return v.Texture()
	}

	return nil
}

//	returns a boolean indicating if the given position is walkable
func (m *Map) Walkable(p *Point) bool {
	if v := m.GetValueAtPoint(p); v != nil {
		return v.Walkable()
	}

	return true
}

func (m *Map) width() int {
	l := len(m.Grid)
	if l == 0 {
		return 0
	}
	return len(m.Grid[0])
}

func (m *Map) height() int {
	return len(m.Grid)
}
