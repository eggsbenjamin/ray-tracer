package game

type Level struct {
	StartingPos           *Point
	Floor, Walls, Ceiling *Map
}

//	returns the width and height
func (l *Level) GetSize() (int, int) {
	return l.Floor.GetSize()
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

func NewLevel(sp *Point, f, w, c *Map) *Level {
	return &Level{
		StartingPos: sp,
		Floor:       f,
		Walls:       w,
		Ceiling:     c,
	}
}
