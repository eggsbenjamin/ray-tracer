package game

type Tile interface {
	Texture() *Texture
	Walkable() bool
	Visable() bool
}

type Interactable interface {
	Interact(*Player, *Level) bool
}

type Wall struct {
	tex *Texture
}

func NewWall(tex *Texture) *Wall {
	return &Wall{
		tex: tex,
	}
}

func (w *Wall) Texture() *Texture {
	return w.tex
}

func (w *Wall) Walkable() bool {
	return false
}

func (w *Wall) Visable() bool {
	return true
}

type Door struct {
	tex  *Texture
	open bool
}

func NewDoor(tex *Texture) *Door {
	return &Door{
		tex: tex,
	}
}

func (d *Door) Texture() *Texture {
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
