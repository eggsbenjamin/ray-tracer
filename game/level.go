package game

import "github.com/veandco/go-sdl2/sdl"

type Level struct {
	StartingPos           *Point
	Floor, Walls, Ceiling *Map
	SkyBox                *sdl.Texture
}

//	returns the width and height
func (l *Level) GetSize() (int, int) {
	return l.Walls.GetSize()
}

//	returns a boolean indicating if the given position is walkable
func (l *Level) Walkable(p *Point) bool {
	return l.Walls.Walkable(p)
}

func (l *Level) Interact(p *Point, pl *Player) bool {
	if i, ok := l.Walls.GetValueAtPoint(p).(Interactable); ok {
		return i.Interact(pl, l)
	}

	return false
}

func NewLevel(sp *Point, f, w, c *Map, sb *sdl.Texture) *Level {
	return &Level{
		StartingPos: sp,
		Floor:       f,
		Walls:       w,
		Ceiling:     c,
		SkyBox:      sb,
	}
}
