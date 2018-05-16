package game

import (
	"math"
)

type Map struct {
	Grid           [][]int
	TexturePalette WallTexturePalette
}

//	constructor
func NewMap(w, h int, tp WallTexturePalette) *Map {
	grid := make([][]int, w)
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, h)
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
func (m *Map) GetValueAtPoint(p *Point) int {
	x := int(math.Floor(p.X))
	y := int(math.Floor(p.Y))
	w, h := m.GetSize()
	if x < 0 || y >= w || y < 0 || x >= h { // out of bounds...
		return -1
	}
	return m.Grid[x][y]
}

func (m *Map) GetTextureAtPoint(p *Point) *Texture {
	if v := m.GetValueAtPoint(p); v != -1 {
		return m.TexturePalette[v]
	}

	return nil
}

//	returns a boolean indicating if the given position is walkable
func (m *Map) Walkable(p *Point) bool {
	return m.GetValueAtPoint(p) == 0
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
