package game

import "github.com/veandco/go-sdl2/sdl"

type Tile interface {
	Texture() *sdl.Texture
	Walkable() bool
	Visable() bool
}

type Interactable interface {
	Interact(*Player, *Level) bool
}

type Wall struct {
	tex *sdl.Texture
}

func NewWall(tex *sdl.Texture) *Wall {
	return &Wall{
		tex: tex,
	}
}

func (w *Wall) Texture() *sdl.Texture {
	return w.tex
}

func (w *Wall) Walkable() bool {
	return false
}

func (w *Wall) Visable() bool {
	return true
}

type Door struct {
	tex  *sdl.Texture
	open bool
}

func NewDoor(tex *sdl.Texture) *Door {
	return &Door{
		tex: tex,
	}
}

func (d *Door) Texture() *sdl.Texture {
	return d.tex
}

func (d *Door) Walkable() bool {
	return d.open
}

func (d *Door) Visable() bool {
	return !d.open
}

func (d *Door) Interact(pl *Player, l *Level) bool {
	curTile := l.Walls.GetValueAtPoint(pl.Pos())
	if d == curTile { // don't open and close door if standing on it
		return false
	}

	d.open = !d.open
	return true
}
